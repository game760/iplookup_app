// 基础请求函数
const TIMEOUT = 5000; // 超时时间5秒
const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || '/api/v1'; // 从环境变量获取API地址

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
    
    if (data.code === 1) {
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

// 导出API方法
export default {
  queryBasicIP(ip) {
    return request(`/ip/query?ip=${encodeURIComponent(ip)}`);
  },
  queryDetailIP(ip) {
    return request(`/ip/detail?ip=${encodeURIComponent(ip)}`);
  },
  getMyIP() {
    return request('/ip/my');
  }
};