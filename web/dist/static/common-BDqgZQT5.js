function t(r){if(Array.isArray(r))return r.map(t);if(typeof r=="object"&&r!==null){const n={};for(const e in r)r.hasOwnProperty(e)&&(n[e]=t(r[e]));return n}else return null}export{t as c};
