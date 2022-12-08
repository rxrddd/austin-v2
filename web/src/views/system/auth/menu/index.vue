<template>
  <div class="app-container">
    <div class="filter-container">
      <el-button class="filter-item" type="primary" icon="el-icon-plus" @click="handleCreate(0)" v-if="checkBtnPermission('createMenu')">
        新增
      </el-button>
    </div>
    <el-table :data="list" row-key="id" style="width: 100%">
      <el-table-column label="菜单ID" min-width="180" prop="id" align="center" />
      <el-table-column label="父菜单ID" min-width="180" prop="parentId" align="center" />
      <el-table-column label="菜单名称" min-width="180" prop="title" align="center" />
      <el-table-column label="图标" min-width="180" prop="icon" align="center">
        <template #default="scope">
          <div v-if="scope.row.icon">
            <i :class="scope.row.icon"></i>
            <span style="margin-left:5px;">{{ scope.row.icon }}</span>
          </div>
        </template>
      </el-table-column>
      <el-table-column label="路由名称" min-width="180" prop="name" align="center" />
      <el-table-column label="路由路径" min-width="180" prop="path" align="center" />
      <el-table-column label="是否隐藏" prop="hidden" align="center">
        <template slot-scope="{row}">
          <span>{{ row.hidden | hiddenInfoFilter }}</span>
        </template>
      </el-table-column>
      <el-table-column label="排序" min-width="180" prop="sort" align="center" />
      <el-table-column label="文件路径" min-width="180" prop="component" align="center" />
      <el-table-column label="创建时间" prop="createdAt" min-width="180">
        <template slot-scope="{row}">
          <span>{{ row.createdAt | timeToDay }}</span>
        </template>
      </el-table-column>
      <el-table-column label="更新时间" prop="updatedAt" min-width="180">
        <template slot-scope="{row}">
          <span>{{ row.updatedAt | timeToDay }}</span>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="250" class-name="small-padding fixed-width" fixed="right">
        <template slot-scope="{row,$index}">
          <el-button type="primary" size="mini" @click="handleChildMenu(row)" v-if="checkBtnPermission('createMenu')">
            新增子菜单
          </el-button>
          <el-button type="primary" size="mini" @click="handleUpdate(row)" v-if="checkBtnPermission('updateMenu')">
            编辑
          </el-button>
          <el-button size="mini" type="danger" @click="handleDelete(row)" v-if="checkBtnPermission('deleteMenu')">
            删除
          </el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-dialog :title="textMap[dialogStatus]" :visible.sync="dialogFormVisible">
      <el-form ref="dataForm" :rules="rules" :model="temp" label-position="left" label-width="150px"
        style="width: 550px; margin-left:50px;">
        <el-form-item label="ID" prop="id" v-if="dialogStatus === 'update'">
          <el-input v-model="temp.id" />
        </el-form-item>
        <el-form-item label="父级菜单" prop="parentIds">
          <el-cascader v-model="temp.parentIds" :options="menuOptions" style="width:100%"
            :props="{ checkStrictly: true, label: 'title', value: 'id', emitPath: 'true' }" :show-all-levels="false"
            @change="handleChange">
          </el-cascader>
        </el-form-item>
        <el-form-item label="菜单名称" prop="title">
          <el-input v-model="temp.title" />
        </el-form-item>
        <el-form-item label="路由路径" prop="path">
          <el-input v-model="temp.path" placeholder="/xxx/xxx"/>
        </el-form-item>
        <el-form-item label="路由name" prop="name">
          <el-input v-model="temp.name" />
        </el-form-item>
        <el-form-item label="是否隐藏" prop="hidden">
          <el-radio v-model="temp.hidden" label="1">是</el-radio>
          <el-radio v-model="temp.hidden" label="0">否</el-radio>
        </el-form-item>
        <el-form-item label="前端文件路径" prop="component">
          <el-input v-model="temp.component" placeholder="/xxx/xxx/index"/>
        </el-form-item>
        <el-form-item label="排序" prop="sort">
          <el-input v-model="temp.sort" />
        </el-form-item>
        <el-form-item label="图标" prop="icon">
          <el-select placeholder="请选择" v-model="temp.icon" filterable>
            <el-option v-for=" item in elementIcons" :label="'el-icon-' + item" :value="'el-icon-' + item">
              <i :class="'el-icon-' + item"></i>
              {{ item }}
            </el-option>
          </el-select>
        </el-form-item>
        <el-button style="margin-top:12px" size="small" type="primary" icon="el-icon-plus"
          @click="addBtn()">新增可控按钮</el-button>
        <el-form-item v-for="(item, i) in menuBtns" :key="item.id"  :label="'按钮' + (i + 1)"  :rules="{required: true, message: '按钮不能为空', trigger: 'blur'}">
          <el-input v-model="item.name" style="width:150px; float: left; " placeholder="请输入按钮名称"/>
          <el-input v-model="item.description" style="width:150px; float: left; margin-left: 15px;" placeholder="请输入按钮描述"/>
          <el-button type="danger" style="float: left; margin-left: 15px;" @click="deleteBtn(i)">删除</el-button>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">
          取消
        </el-button>
        <el-button type="primary" @click="dialogStatus === 'create' ? createData() : updateData()">
          确认
        </el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import { getBaseMenuTree, createMenu, updateMenu, deleteMenu } from '@/api/auth/menu'
import waves from '@/directive/waves' // waves directive
import elementIcons from './element-icons'
import checkBtnPermission from '@/utils/permission'

export default {
  name: 'ComplexTable',
  directives: { waves },
  filters: {
    hiddenInfoFilter(hidden) {
      const hiddenMap = {
        1: '隐藏',
        0: '显示'
      }
      return hiddenMap[hidden]
    },
    timeToDay(times) {
      return times.slice(0, 10)
    },
  },
  watch: {
    filterMenuText(val) {
      this.$refs.menuData.filter(val);
    },
    filterApiText(val) {
      this.$refs.apiData.filter(val);
    },
  },
  data() {
    return {
      elementIcons,
      menuOptions: [
        {
          id: 0,
          title: '根菜单',
        }
      ],
      list: [],
      menuBtns: [],
      temp: {
        id: undefined,
        parent_id: undefined,
        path: undefined,
        name: undefined,
        hidden: undefined,
        component: undefined,
        sort: undefined,
        title: undefined,
        icon: undefined,
        redirect: undefined,
        parentIds:[],
        menuBtns: [],
      },
      dialogFormVisible: false,
      dialogStatus: '',
      textMap: {
        update: '编辑',
        create: '创建'
      },
      rules: {
        title: [{ required: true, message: '菜单名称不得为空', trigger: 'blur' }],
        path: [{ required: true, message: '路由路径不得为空', trigger: 'blur' }],
        name: [{ required: true, message: '路由name不得为空', trigger: 'blur' }],
        hidden: [{ required: true, message: '是否隐藏不得为空', trigger: 'blur' }],
        component: [{ required: true, message: '前端文件路径不得为空', trigger: 'blur' }],
        sort: [{ required: true, message: '排序不得为空', trigger: 'blur' }],
        icon: [{ required: true, message: 'icon不得为空', trigger: 'blur' }],
      },
    }
  },
  created() {
    this.getList()
  },
  methods: {
    checkBtnPermission,
    getList() {
      this.listLoading = true
      getBaseMenuTree().then(response => {
        this.list = response.data.list
      })
    },
    resetTemp() {
      this.temp = {
        id: undefined,
        parent_id: undefined,
        path: undefined,
        name: undefined,
        hidden: undefined,
        component: undefined,
        sort: undefined,
        title: undefined,
        icon: undefined,
        redirect: undefined,
        parentIds:[],
        menuBtns: [],
      }
      this.menuBtns = []
    },
    handleChange(value) {
      this.temp.parentIds = value
    },
    handleCreate(parentId) {
      this.resetTemp()
      this.dialogStatus = 'create'
      this.dialogFormVisible = true
      if (parentId == 0) {
        this.temp.parentIds = [0]
      }
      this.setOptions()
      this.$nextTick(() => {
        this.$refs['dataForm'].clearValidate()
      })
    },
    createData() {
      this.temp.parentId = this.temp.parentIds[this.temp.parentIds.length - 1]
      this.temp.menuBtns = this.menuBtns
      this.$refs['dataForm'].validate((valid) => {
        if (valid) {
          createMenu(this.temp).then(response => {
            this.getList()
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
      this.dialogStatus = 'update'
      this.dialogFormVisible = true
      if (row.parentId == "0") {
        this.temp.parentIds = [0]
      }
      this.setOptions()
      this.menuBtns = row.menuBtns
      this.$nextTick(() => {
        this.$refs['dataForm'].clearValidate()
      })
    },
    updateData() {
      this.temp.parentId = this.temp.parentIds[this.temp.parentIds.length - 1]
      this.temp.menuBtns = this.menuBtns
      this.$refs['dataForm'].validate((valid) => {
        if (valid) {
          const tempData = Object.assign({}, this.temp)
          updateMenu(tempData).then(() => {
            this.getList()
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
    handleDelete(row) {
      deleteMenu({ id: row.id }).then(() => {
        this.getList()
        this.$notify({
          title: 'Success',
          message: '删除成功',
          type: 'success',
          duration: 2000
        })
      })
    },
    setOptions() {
      this.menuOptions = [
        {
          id: 0,
          title: '根菜单'
        }
      ]
      this.setMenuOptions(this.list, this.menuOptions)
    },
    setMenuOptions(dataList, optionsData) {
      this.temp.id = String(this.temp.id)
      dataList &&
        dataList.forEach(item => {
          if (item.children && item.children.length) {
            const option = {
              id: item.id,
              title: item.title,
              children: []
            }
            this.setMenuOptions(
              item.children,
              option.children,
            )
            optionsData.push(option)
          } else {
            const option = {
              id: item.id,
              title: item.title,
            }
            optionsData.push(option)
          }
        })
    },
    handleChange(value) {
    },
    // 新增按钮
    addBtn() {
      if (!this.menuBtns) {
        this.menuBtns = []
      }
      this.menuBtns.push({
        name: '',
        description: '',
      })
    },
    // 删除按钮
    deleteBtn(index) {
      this.menuBtns.splice(index, 1)
    },
    handleChildMenu(row) {
      this.resetTemp()
      this.dialogStatus = 'create'
      this.dialogFormVisible = true
      // 删除首个根元素
      var tmpParentIds = row.parentIds
      if (row.parentIds[0] == "0") {
        tmpParentIds.pop()
      }
      tmpParentIds.push(row.id + '')
      this.temp.parentIds = tmpParentIds
      this.setOptions()
      this.$nextTick(() => {
        this.$refs['dataForm'].clearValidate()
      })
    },
  }
}
</script>
