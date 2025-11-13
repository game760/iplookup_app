<template>
  <div class="container mt-5">
    <div class="row justify-content-center">
      <div class="col-md-8">
        <div class="card shadow-sm">
          <div class="card-header bg-primary text-white">
            <h3 class="mb-0">IP查询工具</h3>
          </div>
          <div class="card-body">
            <div class="input-group mb-3">
              <input
                type="text"
                v-model="ipAddress"
                class="form-control"
                placeholder="请输入IPv4或IPv6地址"
                @keyup.enter="handleQuery"
              >
              <button
                class="btn btn-primary"
                @click="handleQuery"
                :disabled="isQuerying"
              >
                <font-awesome-icon icon="search" class="me-1"></font-awesome-icon>
                查询
              </button>
            </div>
            
            <div class="d-flex gap-2 mb-3">
              <button
                class="btn btn-outline-secondary btn-sm px-4"
                @click="queryMyIp"
                :disabled="isQuerying"
              >
                <font-awesome-icon icon="location-arrow" class="me-1"></font-awesome-icon>
                查询我的IP
              </button>
              <template v-if="recentQueries.length">
                <button 
                  class="btn btn-outline-secondary btn-sm px-4"
                  @click="clearHistory"
                >
                  <font-awesome-icon icon="trash" class="me-1"></font-awesome-icon>
                  清除历史
                </button>
              </template>
            </div>
            
            <!-- 错误提示 -->
            <div v-if="error" class="invalid-feedback d-block mt-2 animate-fade-in">
              <font-awesome-icon icon="exclamation-circle" class="me-1"></font-awesome-icon>
              {{ error }}
            </div>

            <!-- 加载状态 -->
            <div v-if="loading" class="text-center my-8">
              <div class="spinner-border text-primary" role="status">
                <span class="visually-hidden">Loading...</span>
              </div>
              <p class="mt-3 text-muted">正在查询，请稍候...</p>
            </div>
            
            <!-- 查询结果 -->
            <div v-if="result && !loading" class="mt-4">
              <div v-if="result.data" class="card">
                <div class="card-header">
                  <h5 class="mb-0">
                    <template v-if="result.data.IP">
                      {{ result.data.IP }} 
                      <span class="badge bg-secondary" v-if="result.data.CountryName">
                        {{ result.data.CountryName }}
                      </span>
                      <span class="badge bg-info ms-1">
                        {{ isIPv6(result.data.IP) ? 'IPv6' : 'IPv4' }}
                      </span>
                    </template>
                    <template v-else>
                      IP查询结果
                    </template>
                  </h5>
                </div>
                <div class="card-body">
                  <div class="row">
                    <div class="col-md-6 mb-3">
                      <strong>国家/地区:</strong> {{ result.data.CountryName || '未知' }}
                    </div>
                    <div class="col-md-6 mb-3">
                      <strong>区域:</strong> {{ result.data.Region || '未知' }}
                    </div>
                    <div class="col-md-6 mb-3">
                      <strong>省份:</strong> {{ result.data.Province || '未知' }}
                    </div>
                    <div class="col-md-6 mb-3">
                      <strong>城市:</strong> {{ result.data.City || '未知' }}
                    </div>
                    <div class="col-md-6 mb-3">
                      <strong>运营商:</strong> {{ result.data.ISP || '未知' }}
                    </div>
                    <div class="col-md-6 mb-3">
                      <strong>坐标:</strong> {{ result.data.Latitude }}, {{ result.data.Longitude }}
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
        
        <!-- 历史记录 -->
        <div v-if="recentQueries.length" class="mt-4 card">
          <div class="card-header">
            <h5 class="mb-0">查询历史</h5>
          </div>
          <div class="card-body">
            <ul class="list-group">
              <li 
                v-for="(item, index) in recentQueries" 
                :key="index" 
                class="list-group-item d-flex justify-content-between align-items-center"
              >
                <div>
                  <strong>{{ item.ip }}</strong>
                  <small class="text-muted ms-2">{{ item.type }}</small>
                </div>
                <button 
                  class="btn btn-sm btn-outline-primary"
                  @click="requery(item.ip)"
                >
                  再次查询
                </button>
              </li>
            </ul>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import ipService from '../services/ipservice';

// 状态管理
const ipAddress = ref('');
const result = ref(null);
const error = ref('');
const loading = ref(false);
const isQuerying = ref(false);
const recentQueries = ref([]);

// 判断是否为IPv6地址
const isIPv6 = (ip) => {
  return ip.includes(':');
};

// 自动识别IP类型并查询
const handleQuery = async () => {
  if (!ipAddress.value.trim()) {
    error.value = '请输入IP地址';
    return;
  }

  loading.value = true;
  isQuerying.value = true;
  error.value = '';
  
  try {
    // 调用自动识别接口
    const data = await ipService.queryIP(ipAddress.value.trim());
    result.value = data;
    
    // 记录查询历史
    if (data.data?.IP) {
      recentQueries.value.unshift({
        ip: data.data.IP,
        type: isIPv6(data.data.IP) ? 'IPv6' : 'IPv4'
      });
      // 限制历史记录长度
      if (recentQueries.value.length > 10) {
        recentQueries.value = recentQueries.value.slice(0, 10);
      }
    }
  } catch (err) {
    error.value = err.message || '查询失败，请重试';
  } finally {
    loading.value = false;
    isQuerying.value = false;
  }
};

// 重新查询历史记录中的IP
const requery = async (ip) => {
  ipAddress.value = ip;
  await handleQuery();
};

// 查询本机IP
const queryMyIp = async () => {
  loading.value = true;
  isQuerying.value = true;
  error.value = '';
  
  try {
    const data = await ipService.getMyIP();
    result.value = data;
    
    // 同步更新输入框和历史记录
    if (data.data?.IP) {
      ipAddress.value = data.data.IP;
      recentQueries.value.unshift({
        ip: data.data.IP,
        type: isIPv6(data.data.IP) ? 'IPv6' : 'IPv4'
      });
      
      if (recentQueries.value.length > 10) {
        recentQueries.value = recentQueries.value.slice(0, 10);
      }
    }
  } catch (err) {
    error.value = err.message || '查询失败，请重试';
  } finally {
    loading.value = false;
    isQuerying.value = false;
  }
};

// 清除历史记录
const clearHistory = () => {
  recentQueries.value = [];
};
</script>