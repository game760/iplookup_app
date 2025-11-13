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
                class="btn btn-outline-primary btn-sm px-4"
                @click="handleIPv4Query"
                :disabled="isQuerying"
              >
                IPv4查询
              </button>
              <button
                class="btn btn-outline-primary btn-sm px-4"
                @click="handleIPv6Query"
                :disabled="isQuerying"
              >
                IPv6查询
              </button>
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
                      <span class="badge bg-secondary" v-if="result.data.CountryCode">
                        {{ result.data.CountryCode }}
                      </span>
                    </template>
                    <template v-else>
                      IP查询结果
                    </template>
                  </h5>
                </div>
                <div class="card-body">
                  <div class="row">
                    <!-- IPv4结果展示 -->
                    <template v-if="result.data.Domain">
                      <div class="col-md-6 mb-3">
                        <strong>国家:</strong> {{ result.data.CountryName }}
                      </div>
                      <div class="col-md-6 mb-3">
                        <strong>地区:</strong> {{ result.data.Region }}
                      </div>
                      <div class="col-md-6 mb-3">
                        <strong>城市:</strong> {{ result.data.City }}
                      </div>
                      <div class="col-md-6 mb-3">
                        <strong>时区:</strong> {{ result.data.TimeZone }}
                      </div>
                      <div class="col-md-6 mb-3">
                        <strong>ISP:</strong> {{ result.data.ISP }}
                      </div>
                      <div class="col-md-6 mb-3">
                        <strong>域名:</strong> {{ result.data.Domain }}
                      </div>
                      <div class="col-md-6 mb-3">
                        <strong>使用类型:</strong> {{ result.data.UsageType }}
                      </div>
                      <div class="col-md-6 mb-3">
                        <strong>ASN:</strong> {{ result.data.ASN }}
                      </div>
                      <div class="col-md-6 mb-3">
                        <strong>AS名称:</strong> {{ result.data.ASName }}
                      </div>
                      <div class="col-md-6 mb-3">
                        <strong>坐标:</strong> {{ result.data.Latitude }}, {{ result.data.Longitude }}
                      </div>
                    </template>
                    
                    <!-- IPv6结果展示 -->
                    <template v-else-if="result.data.Network">
                      <div class="col-md-6 mb-3">
                        <strong>国家:</strong> {{ result.data.CountryName }}
                      </div>
                      <div class="col-md-6 mb-3">
                        <strong>地区:</strong> {{ result.data.Region }}
                      </div>
                      <div class="col-md-6 mb-3">
                        <strong>城市:</strong> {{ result.data.City }}
                      </div>
                      <div class="col-md-6 mb-3">
                        <strong>时区:</strong> {{ result.data.TimeZone }}
                      </div>
                      <div class="col-md-6 mb-3">
                        <strong>ISP:</strong> {{ result.data.ISP }}
                      </div>
                      <div class="col-md-6 mb-3">
                        <strong>ASN:</strong> {{ result.data.ASN }}
                      </div>
                      <div class="col-md-6 mb-3">
                        <strong>AS名称:</strong> {{ result.data.ASName }}
                      </div>
                      <div class="col-md-6 mb-3">
                        <strong>网络:</strong> {{ result.data.Network }}
                      </div>
                      <div class="col-md-6 mb-3">
                        <strong>坐标:</strong> {{ result.data.Latitude }}, {{ result.data.Longitude }}
                      </div>
                    </template>
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
                  @click="requery(item.ip, item.method)"
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