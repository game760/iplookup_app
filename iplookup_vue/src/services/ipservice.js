// 从环境变量获取API地址
const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || '/api';
const TIMEOUT = 10000; // 10秒超时

const request = async (url) => {
  try {
    const controller = new AbortController();
    const timeoutId = setTimeout(() => controller.abort(), TIMEOUT);
    
    const response = await fetch(`${API_BASE_URL}${url}`, {
      signal: controller.signal,
      headers: {
        'Accept': 'application/json'
      }
    });
    
    clearTimeout(timeoutId);
    
    const data = await response.json();
    
    if (!response.ok) {
      throw new Error(data.message || '请求失败');
    }
    
    if (data.code !== 0) {
      throw new Error(data.message || '服务器返回错误');
    }
    
    return data;
  } catch (error) {
    if (error.name === 'AbortError') {
      throw new Error(`请求超时（${TIMEOUT}ms），请重试`);
    }
    throw error;
  }
};

// 自动识别IP类型并查询
const queryIP = (ip) => {
  return request(`/ip/query?ip=${encodeURIComponent(ip)}`);
};

// 查询本机IP
const getMyIP = () => {
  return request('/ip/my');
};

export default {
  queryIP,
  getMyIP
};