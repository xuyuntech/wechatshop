/**
 * @author cabernety
 */

/* eslint-disable */
import 'isomorphic-fetch';
import Cookies from 'universal-cookie';
import { isString, isFunction } from 'lodash';

const logger = console;
const isNode =
  Object.prototype.toString.call(
    typeof process !== 'undefined' ? process : 0,
  ) === '[object process]';

const baseURL = isNode ? process.env.API_SERVER_URL : window.App.apiUrl;

/* eslint-disable no-console,no-undef */
const cookies = new Cookies();

const STATUS = {
  OK: 0,
};

class Err {
  type = 'ErrDefault';
  options = {
    message: 'default err',
  };
  constructor(options) {
    let opt = options;
    if (isString(options)) {
      opt = {
        message: options,
      };
    }
    this.options = opt;
  }
  is(err) {
    if (!err) {
      return false;
    }
    return this instanceof err;
  }
  toString() {
    const { message } = this.options;
    return `[BFetch Err] -> ${this.type}:${message}`;
  }
}

export class ErrAPI extends Err {
  type = 'ErrAPI';
}

export class ErrJSONParse extends Err {
  type = 'ErrJSONParse';
}

export class ErrRequest extends Err {
  type = 'ErrRequest';
}

const addParams = (url, params) => {
  if (!params) {
    return url;
  }
  const arr = Object.keys(params).map(key => {
    let param = params[key];
    if (typeof param === 'undefined' || isFunction(param)) {
      param = '';
    }
    return `${key}=${encodeURIComponent(param)}`;
  });
  if (!arr.length) {
    return url;
  }
  return `${url}${/\?/.test(url) ? '&' : '?'}${arr.join('&')}`;
};

const bFetch = async (url, options) => {
  const defaults = {
    method: 'GET',
    headers: {
      Accept: 'application/json',
      'Content-Type': 'application/json',
      'X-Access-Token': cookies.get('X-Access-Token'),
    },
  };
  const params = {
    ...options.params,
    ...options.pagination,
  };
  try {
    const uUrl = addParams(`${baseURL}${url}`, params);
    logger.log(`fetch: ${uUrl}`);
    const res = await fetch(uUrl, {
      ...defaults,
      ...options,
      headers: {
        ...defaults.headers,
        ...(options && options.headers),
        ...(options.token && { 'X-Access-Token': options.token }),
      },
    });
    switch (res.status) {
      case 200: {
        const txt = await res.text();
        let json = null;
        try {
          json = JSON.parse(txt);
        } catch (err) {
          throw new ErrJSONParse(`JSON 解析失败: ${txt}`);
        }
        if (json.status === STATUS.OK) {
          return json;
        }
        throw new ErrAPI(`api 返回错误 status:${json.status}`);
      }
      default:
        throw new ErrRequest(`请求失败 ${res.status}:${res.statusText}`);
    }
  } catch (err) {
    logger.error(err);
    throw err;
  }
};

export default bFetch;
