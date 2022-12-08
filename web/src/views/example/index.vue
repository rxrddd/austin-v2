<template>
  <div class="app-container">
    <div class="filter-container">
      <div>
        测试按钮权限:
        <el-button size="mini" type="success" v-if="checkBtnPermission('hasPermissionButton')">
          有权限
        </el-button>
        无权限
        <el-button size="mini" type="success" v-if="checkBtnPermission('noPermissionButton')">
          无权限
        </el-button>
      </div>
    </div>
  </div>
</template>

<script>
import checkBtnPermission from '@/utils/permission'

export default {
  name: 'test',
  data() {
    return {
      
    }
  },
  created() {

  },
  methods: {
    checkBtnPermission,
  }
}
</script>

