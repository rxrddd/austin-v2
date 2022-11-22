<template>
  <div class="app-container">
    <div class="filter-container">
      <el-button class="filter-item" type="primary" icon="el-icon-edit" @click="handleCreate">
        新增
      </el-button>

      <el-input v-model="listQuery.mobile" placeholder="手机号" style="width: 200px;" class="filter-item" @keyup.enter.native="handleFilter" />
      
      <el-input v-model="listQuery.username" placeholder="用户名" style="width: 200px;" class="filter-item" @keyup.enter.native="handleFilter" />
      
      <el-select v-model="listQuery.status" placeholder="状态" clearable class="filter-item" style="width: 130px">
        <el-option v-for="item in statusOptions" :key="item.key" :label="item.display_name" :value="item.key" />
      </el-select>
      
      <el-date-picker
      v-model="createdSearch"
      type="daterange"
      range-separator="至"
      start-placeholder="创建开始日期"
      end-placeholder="创建结束日期">
    </el-date-picker>

      <el-button v-waves class="filter-item" type="primary" icon="el-icon-search" @click="handleFilter">
        搜索
      </el-button>
      
    </div>

    <el-table
      v-loading="listLoading"
      :data="list"
      border 
      style="width: 100%;"
    >
      <el-table-column label="ID" prop="id" align="center" width="50">
        <template slot-scope="{row}">
          <span>{{ row.id }}</span>
        </template>
      </el-table-column>
      <el-table-column label="用户名" prop="username" align="center">
        <template slot-scope="{row}">
          <span>{{ row.username }}</span>
        </template>
      </el-table-column>
      <el-table-column label="角色" prop="role" align="center">
        <template slot-scope="{row}">
          <span>{{ row.role }}</span>
        </template>
      </el-table-column>
      <el-table-column label="手机号" prop="mobile" align="center">
        <template slot-scope="{row}">
          <span>{{ row.mobile }}</span>
        </template>
      </el-table-column>
      <el-table-column label="昵称" prop="nickname" align="center">
        <template slot-scope="{row}">
          <span>{{ row.nickname }}</span>
        </template>
      </el-table-column>
      <el-table-column label="状态" class-name="status-col">
        <template slot-scope="{row}">
          <el-tag :type="row.status | statusTagFilter">
            {{ row.status | statusFilter}}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column label="登录时间" prop="lastLoginTime" align="center">
        <template slot-scope="{row}">
          <span>{{ row.lastLoginTime  | timeToDay}}</span>
        </template>
      </el-table-column>
      <el-table-column label="登录ip" prop="lastLoginIp" align="center">
        <template slot-scope="{row}">
          <span>{{ row.lastLoginIp }}</span>
        </template>
      </el-table-column>
      <el-table-column label="创建时间" prop="createdAt" align="center">
        <template slot-scope="{row}">
          <span>{{ row.createdAt | timeToDay}}</span>
        </template>
      </el-table-column>
      <el-table-column label="更新时间" prop="updatedAt" align="center">
        <template slot-scope="{row}">
          <span>{{ row.updatedAt | timeToDay}}</span>
        </template>
      </el-table-column>
      <el-table-column label="删除时间" prop="updatedAt" align="center">
        <template slot-scope="{row}">
          <span>{{ row.deletedAt | timeToDay}}</span>
        </template>
      </el-table-column>
      <el-table-column label="操作" align="center" width="250" class-name="small-padding fixed-width">
        <template slot-scope="{row,$index}">
          <el-button type="primary" size="mini" @click="handleUpdate(row)">
            编辑
          </el-button>
          <el-button v-if="row.deletedAt == ''" size="mini" type="danger" @click="handleDelete(row,$index)">
            删除
          </el-button>
          <el-button v-if="row.deletedAt != ''" size="mini" type="success" @click="handleRecover(row,$index)">
            恢复
          </el-button>
          <el-button v-if="row.status == '1'" size="mini" type="warning" @click="handleForbid(row,$index)">
            禁用
          </el-button>
          <el-button v-if="row.status == '2'" size="mini" type="success" @click="handleApprove(row,$index)">
            解禁
          </el-button>
        </template>
      </el-table-column>
    </el-table>

    <pagination v-show="total>0" :total="total" :page.sync="listQuery.page" :limit.sync="listQuery.limit" @pagination="getList" />

    <el-dialog :title="textMap[dialogStatus]" :visible.sync="dialogFormVisible">
      <el-form ref="dataForm" :rules="rules" :model="temp" label-position="left" label-width="70px" style="width: 400px; margin-left:50px;">
        <el-form-item label="ID" prop="id" v-if="temp.id != undefined">
          <el-input v-model="temp.id"/>
        </el-form-item>
        <el-form-item label="用户名" prop="username">
          <el-input v-model="temp.username" :disabled="isDisabled"/>
        </el-form-item>
        <el-form-item label="密码" prop="password">
          <el-input v-model="temp.password"/>
        </el-form-item>
        <el-form-item label="昵称" prop="nickname">
          <el-input v-model="temp.nickname" />
        </el-form-item>
        <el-form-item label="手机号" prop="mobile">
          <el-input v-model="temp.mobile" />
        </el-form-item>
        <el-form-item label="头像" prop="avatar">
          <el-input v-model="temp.avatar" />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="temp.status" class="filter-item" placeholder="Please select">
            <el-option v-for="item in statusOptions" :key="item.key" :label="item.display_name" :value="item.key" />
          </el-select>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">
          取消
        </el-button>
        <el-button type="primary" @click="dialogStatus==='create'?createData():updateData()">
          确认
        </el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import { listAdministrator, createAdministrator, updateAdministrator, deleteAdministrator, recoverAdministrator, forbidAdministrator, approveAdministrator} from '@/api/administrator'
import waves from '@/directive/waves' // waves directive
import Pagination from '@/components/Pagination' // secondary package based on el-pagination
import { getDate, parseTime } from '@/utils/index.js'
const statusOptions = [
  { key: '1', display_name: '正常' },
  { key: '2', display_name: '禁用' },
]

// arr to obj, such as { CN : "China", US : "USA" }
const calendarTypeKeyValue = statusOptions.reduce((acc, cur) => {
  acc[cur.key] = cur.display_name
  return acc
}, {})

export default {
  name: 'ComplexTable',
  components: { Pagination },
  directives: { waves },
  filters: {
    statusFilter(status) {
      const statusMap = {
        1: '正常',
        2: '禁用'
      }
      return statusMap[status]
    },
    statusTagFilter(status){
      const statusMap = {
        1: 'success',
        2: 'danger'
      }
      return statusMap[status]
    },
    timeToDay(times) {
      return times.slice(0, 10)
    },
  },
  data() {
    return {
      isDisabled:false,
      createdSearch:'',
      statusOptions,
      list: null,
      total: 0,
      listLoading: true,
      listQuery: {
        page: 1,
        limit: 20,
        mobile: undefined,
        username: undefined,
        status: undefined,
        created_at_start: undefined,
        created_at_end: undefined,
      },
      temp: {
        id: undefined,
        status: undefined,
        username: undefined,
        password: undefined,
        mobile: undefined,
        nickname: undefined,
        avatar: undefined,
        status: undefined,
      },
      dialogFormVisible: false,
      dialogStatus: '',
      textMap: {
        update: '编辑',
        create: '创建'
      },
      dialogPvVisible: false,
      pvData: [],
      rules: {
        username: [{ required: true, message: '用户名不得为空', trigger: 'blur' }],
        nickname: [{ required: true, message: '昵称不得为空', trigger: 'blur' }],
        password: [{ required: true, message: '密码不得为空', trigger: 'change' }],
        mobile: [{ required: true, message: '手机号不得为空', trigger: 'blur' }],
        avatar: [{ required: true, message: '头像不得为空', trigger: 'blur' }],
      },
      downloadLoading: false
    }
  },
  created() {
    this.getList()
  },
  methods: {
    getList() {
      this.listLoading = true
      listAdministrator(this.listQuery).then(response => {
        this.list = response.data.list
        this.total = parseInt(response.data.total)
        this.listLoading = false
      })
    },
    handleFilter() {
      if(this.createdSearch != ''){
        this.listQuery.created_at_start = parseTime(this.createdSearch[0], "{y}-{m}-{d}")
        this.listQuery.created_at_end = parseTime(this.createdSearch[1], "{y}-{m}-{d}")
      }else{
        this.listQuery.created_at_start = ""
        this.listQuery.created_at_end = ""
      }
      this.listQuery.page = 1
      this.getList()
    },
    handleModifyStatus(row, status) {
      this.$message({
        message: '操作Success',
        type: 'success'
      })
      row.status = status
    },
    resetTemp() {
      this.isDisabled = false;
      this.temp = {
        id: undefined,
        status: undefined,
        username: undefined,
        password: undefined,
        mobile: undefined,
        nickname: undefined,
        avatar: undefined,
        status: undefined,
      }
    },
    handleCreate() {
      this.resetTemp()
      this.dialogStatus = 'create'
      this.dialogFormVisible = true
      this.$nextTick(() => {
        this.$refs['dataForm'].clearValidate()
      })
    },
    createData() {
      this.$refs['dataForm'].validate((valid) => {
        if (valid) {
          createAdministrator(this.temp).then(response => {
            this.list.push(response.data)
            this.dialogFormVisible = false
            this.$notify({
              title: 'Success',
              message: '创建成功',
              type: 'success',
              duration: 2000
            })
          })
        }
      })
    },
    handleUpdate(row) {
      this.temp = Object.assign({}, row) // copy obj
      this.isDisabled = true;
      this.dialogStatus = 'update'
      this.dialogFormVisible = true
      this.$nextTick(() => {
        this.$refs['dataForm'].clearValidate()
      })
    },
    updateData() {
      this.$refs['dataForm'].validate((valid) => {
        if (valid) {
          const tempData = Object.assign({}, this.temp)
          updateAdministrator(tempData).then(() => {
            const index = this.list.findIndex(v => v.id === this.temp.id)
            this.list.splice(index, 1, this.temp)
            this.dialogFormVisible = false
            this.$notify({
              title: 'Success',
              message: '更新成功',
              type: 'success',
              duration: 2000
            })
          })
        }
      })
    },
   
    handleDelete(row, index) {
      deleteAdministrator({id:row.id}).then(() => {
        this.list[index].deletedAt = getDate()
        this.$notify({
          title: 'Success',
          message: '删除成功',
          type: 'success',
          duration: 2000
        })
      })
    },
    handleRecover(row, index) {
      recoverAdministrator({id:row.id}).then(() => {
        this.list[index].deletedAt = ""
        this.$notify({
          title: 'Success',
          message: '恢复成功',
          type: 'success',
          duration: 2000
        })
      })
    },
    handleForbid(row, index) {
      forbidAdministrator({id:row.id}).then(() => {
        this.list[index].status = 2
        this.$notify({
          title: 'Success',
          message: '恢复成功',
          type: 'success',
          duration: 2000
        })
      })
    },
    handleApprove(row, index) {
      approveAdministrator({id:row.id}).then(() => {
        this.list[index].status = 1
        this.$notify({
          title: 'Success',
          message: '恢复成功',
          type: 'success',
          duration: 2000
        })
      })
    },
  }
}
</script>
