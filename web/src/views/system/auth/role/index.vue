<template>
  <div class="app-container">
    <div class="filter-container">
      <el-button class="filter-item" type="primary" icon="el-icon-plus" @click="handleCreate(0)">
        新增
      </el-button>
    </div>
    <el-table :data="list" :props="roleTreeProps" row-key="id" style="width: 100%">
      <el-table-column label="角色ID" min-width="180" prop="id" />
      <el-table-column align="center" label="角色名称" min-width="180" prop="name" />
      <el-table-column label="创建时间" prop="createdAt" align="center">
        <template slot-scope="{row}">
          <span>{{ row.createdAt | timeToDay }}</span>
        </template>
      </el-table-column>
      <el-table-column label="更新时间" prop="updatedAt" align="center">
        <template slot-scope="{row}">
          <span>{{ row.updatedAt | timeToDay }}</span>
        </template>
      </el-table-column>
      <el-table-column label="操作" align="center" width="500" class-name="small-padding fixed-width">
        <template slot-scope="{row,$index}">
          <el-button type="warning" size="mini" @click="openDrawer(row)">
            设置权限
          </el-button>
          <el-button type="primary" size="mini" @click="handleChildRole(row)">
            新增子角色
          </el-button>
          <el-button type="primary" size="mini" @click="handleUpdate(row)">
            编辑
          </el-button>
          <el-button size="mini" type="danger" @click="handleDelete(row, $index)">
            删除
          </el-button>
          
        </template>
      </el-table-column>
    </el-table>

    <el-dialog :title="textMap[dialogStatus]" :visible.sync="dialogFormVisible">
      <el-form ref="dataForm" :rules="rules" :model="temp" label-position="left" label-width="100px"
        style="width: 500px; margin-left:50px;">
        <el-form-item label="ID" prop="id" v-if="dialogStatus === 'update'">
          <el-input v-model="temp.id" />
        </el-form-item>
        <el-form-item label="父级角色" prop="parentIds">
          <el-cascader v-model="temp.parentIds" :options="roleOptions" style="width:100%"
            :props="{ checkStrictly: true, label: 'name', value: 'id', emitPath: 'true' }" :show-all-levels="false"
            @change="handleChange">
          </el-cascader>
        </el-form-item>
        <el-form-item label="角色名称" prop="name">
          <el-input v-model="temp.name" />
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
    <el-drawer title="角色配置" :visible.sync="drawer" :with-header="false" size="40%">
      <el-tabs type="border-card">
        <el-tab-pane label="角色菜单">
          <div>
            <el-input v-model="filterMenuText" style="width: 60%;" placeholder="筛选" />
            <el-button style="float: right;" size="small" type="primary" @click="saveMenu">确 定</el-button>
          </div>
          <div class="tree-content">
            <el-tree :data="menuData" show-checkbox node-key="id" default-expand-all highlight-current ref="menuData"
              :default-checked-keys="menuCheckedIds" :props="menuDataProps" :filter-node-method="filterMenuNode">
              <span class="custom-tree-node" slot-scope="{ node, menuData }">
                <span>{{ node.label }}</span>
                <el-button type="text" style="margin-left:10px" v-if="(node.data.menuBtns.length > 0)" size="small"
                  @click="() => assignBtn(node)">
                  分配按钮
                </el-button>
              </span>
            </el-tree>
          </div>
        </el-tab-pane>
        <el-tab-pane label="角色api">
          <div>
            <el-input v-model="filterApiText" style="width: 60%;" placeholder="筛选" />
            <el-button style="float: right;" size="small" type="primary" @click="saveApi">确 定</el-button>
          </div>
          <div class="tree-content">
            <el-tree :data="apiData" show-checkbox node-key="id" default-expand-all highlight-current ref="apiData"
              :default-checked-keys="apiCheckedIds" :props="apiDataProps" :filter-node-method="filterApiNode">
            </el-tree>
          </div>

        </el-tab-pane>
      </el-tabs>
    </el-drawer>

    <el-dialog :visible.sync="btnVisible" title="分配按钮" destroy-on-close>
      <el-table ref="multipleTable" :data="btnData" row-key="ID" @selection-change="handleSelectChange">
        <el-table-column type="selection" />
        <el-table-column label="按钮名称" prop="name" />
        <el-table-column label="按钮描述" prop="description" />
      </el-table>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="saveBtn">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
import checkBtnPermission from '@/utils/permission' // 按钮权限判断函数
import { listRole, createRole, updateRole, deleteRole, saveRoleMenu, getRolePolicies, saveRolePolicies, getRoleMenuBtn, setRoleMenuBtn } from '@/api/auth/role'
import { getBaseMenuTree, getRoleMenu } from '@/api/auth/menu'
import { getApiAll } from '@/api/auth/api'
import waves from '@/directive/waves' // waves directive

export default {
  name: 'ComplexTable',
  directives: { waves },
  filters: {
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
      btnVisible: false,
      btnData: [],
      currentRow: {},
      currentMenuId: undefined,
      filterMenuText: '',
      tmpMenuCheckedIds: [],
      menuCheckedIds: [],
      menuData: [],
      menuDataProps: {
        children: 'children',
        label: 'title'
      },
      filterApiText: '',
      tmpApiCheckedIds: [],
      apiCheckedIds: [],
      apiData: [],
      apiDataProps: {
        children: 'children',
        label: 'name'
      },
      drawer: false,
      allOptionList: [],
      roleOptions: [
        {
          id: 0,
          name: '根角色',
        }
      ],
      roleTreeProps: {
        children: 'children',
        label: 'name'
      },
      list: [],
      listLoading: true,
      listQuery: {
        page: 1,
        limit: 20,
      },
      temp: {
        id: undefined,
        name: undefined,
        parentId: undefined,
        parentIds: undefined,
      },
      dialogFormVisible: false,
      dialogStatus: '',
      textMap: {
        update: '编辑',
        create: '创建'
      },
      selectedBtnIds: [],
      rules: {
        name: [{ required: true, message: '角色名称不得为空', trigger: 'blur' }],
      },
    }
  },
  created() {
    this.getList()
    this.getMenuData()
    this.getApiData()
  },
  methods: {
    checkBtnPermission,
    handleChange(value) {
      this.temp.parentIds = value
    },
    getList() {
      this.listLoading = true
      listRole(this.listQuery).then(response => {
        this.list = response.data.list
        this.listLoading = false
      })
    },
    getMenuData() {
      getBaseMenuTree().then(response => {
        this.menuData = response.data.list
      })
    },
    getApiData() {
      getApiAll().then(response => {
        this.apiData = this.buildApiTree(response.data.list)
      })
    },
    buildApiTree(apis) {
      const apiObj = {}
      apis &&
        apis.forEach(item => {
          item.id = 'path:' + item.path + '-' + 'method:' + item.method
          if (Object.prototype.hasOwnProperty.call(apiObj, item.group)) {
            apiObj[item.group].push(item)
          } else {
            Object.assign(apiObj, { [item.group]: [item] })
          }
        })
      const apiTree = []
      for (const key in apiObj) {
        const treeNode = {
          id: key,
          name: key,
          children: apiObj[key]
        }
        apiTree.push(treeNode)
      }
      return apiTree
    },
    handleFilter() {
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
      this.temp = {
        id: undefined,
        name: undefined,
        parentId: undefined,
        parentIds: undefined,
      }
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
    handleChildRole(row) {
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
    createData() {
      this.temp.parentId = this.temp.parentIds[this.temp.parentIds.length - 1]
      this.$refs['dataForm'].validate((valid) => {
        if (valid) {
          createRole(this.temp).then(response => {
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
      this.$nextTick(() => {
        this.$refs['dataForm'].clearValidate()
      })
    },
    updateData() {
      this.temp.parentId = this.temp.parentIds[this.temp.parentIds.length - 1]
      this.$refs['dataForm'].validate((valid) => {
        if (valid) {
          const tempData = Object.assign({}, this.temp)
          updateRole(tempData).then(() => {
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

    handleDelete(row, index) {
      deleteRole({ id: row.id }).then(() => {
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
      this.roleOptions = [
        {
          id: 0,
          name: '根角色'
        }
      ]
      this.setRoleOptions(this.list, this.roleOptions)
    },
    setRoleOptions(dataList, optionsData) {
      this.temp.id = String(this.temp.id)
      dataList &&
        dataList.forEach(item => {
          if (item.children && item.children.length) {
            const option = {
              id: item.id,
              name: item.name,
              children: []
            }
            this.setRoleOptions(
              item.children,
              option.children,
            )
            optionsData.push(option)
          } else {
            const option = {
              id: item.id,
              name: item.name,
            }
            optionsData.push(option)
          }
        })
    },
    openDrawer(row) {
      this.currentRow = row
      this.drawer = true
      getRoleMenu({ role: row.name }).then(response => {
        this.$refs.menuData.setCheckedKeys([]);
        this.tmpMenuCheckedIds = [];
        this.setRoleMenuChecked(response.data.list)
        this.menuCheckedIds = this.tmpMenuCheckedIds
      })
      getRolePolicies({ role: row.name }).then(response => {
        const apis = response.data.policyRules
        this.tmpApiCheckedIds = [];
        this.$refs.apiData.setCheckedKeys([]);
        apis.forEach(item => {
          this.tmpApiCheckedIds.push('path:' + item.path + '-' + 'method:' + item.method)
        })
        this.apiCheckedIds = this.tmpApiCheckedIds
        return
      })
    },
    setRoleMenuChecked(menuList) {
      menuList && menuList.forEach(item => {
        if (item.children && item.children.length) {
          this.setRoleMenuChecked(item.children)
        } else {
          this.tmpMenuCheckedIds.push(parseInt(item.id))
        }
      })
    },
    filterMenuNode(value, data) {
      if (!value) return true;
      return data.title.indexOf(value) !== -1;
    },
    filterApiNode(value, data) {
      if (!value) return true;
      return data.name.indexOf(value) !== -1;
    },
    saveMenu() {
      const checkArr = this.$refs.menuData.getCheckedNodes(false, true)
      const menuIds = [];
      checkArr.forEach(item => {
        menuIds.push(item.id)
      })
      saveRoleMenu({
        "role_id": this.currentRow.id,
        "menu_ids": menuIds
      }).then(() => {
        this.$notify({
          title: 'Success',
          message: '保存成功',
          type: 'success',
          duration: 1000
        })
      })
    },
    saveApi() {
      const checkArr = this.$refs.apiData.getCheckedNodes(false, true)

      const policyRules = [];
      checkArr.forEach(item => {
        // 只获取有效api，不获取分组名称
        if (item.createdAt != undefined) {
          policyRules.push({
            "path": item.path,
            "method": item.method,
          })
        }
      })
      saveRolePolicies({
        "role": this.currentRow.name,
        "policyRules": policyRules
      }).then(() => {
        this.$notify({
          title: 'Success',
          message: '保存成功',
          type: 'success',
          duration: 1000
        })
      })
    },
    assignBtn(node) {
      this.currentMenuId = node.data.id
      this.btnVisible = true
      this.btnData = []

      node.data.menuBtns.forEach(menuBtn => {
        this.btnData.push(menuBtn)
      });
      // 获取当前角色拥有的菜单按钮
      getRoleMenuBtn({ role_id: this.currentRow.id, menu_id: node.data.id }).then(response => {
        this.btnData.forEach(item => {
          response.data.menuBtnIds.some(btnId => {
            this.$nextTick(() => {
              if (item.id === btnId) {
                this.$refs.multipleTable.toggleRowSelection(item, true)
              }
            })
          })
        })
      })
    },
    saveBtn() {
      setRoleMenuBtn({ role_id: this.currentRow.id, menu_id: this.currentMenuId, menu_btn_ids: this.selectedBtnIds }).then(() => {
        this.$notify({
          title: 'Success',
          message: '保存成功',
          type: 'success',
          duration: 1000
        })
        this.btnVisible = false
      })
    },
    closeDialog() {
      this.btnVisible = false
      this.btnData = []
    },
    handleSelectChange(val) {
      this.selectedBtnIds = [];
      val.forEach(element => {
        this.selectedBtnIds.push(element.id)
      });
    }
  }
}
</script>
<style lang="scss">
.tree-content {
  overflow: auto;
  height: calc(100vh - 100px);
  margin-top: 10px;
}
</style>
