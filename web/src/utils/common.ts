
export function clearValues(obj:any) :any|null{
    if (Array.isArray(obj)) {
      // 如果是数组，递归处理数组中的每一个元素
      return obj.map(clearValues);
    } else if (typeof obj === 'object' && obj !== null) {
      // 如果是对象，递归处理对象中的每一个属性
      const result = {};
      for (const key in obj) {
        if (obj.hasOwnProperty(key)) {
          result[key] = clearValues(obj[key]);
        }
      }
      return result;
    } else {
      // 基本类型（如字符串、数字等），直接返回 null 或其他默认值
      return null; // 或者使用 '' 空字符串等其他默认值
    }
  }