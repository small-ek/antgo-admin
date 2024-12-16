import {reactive, ref} from "vue";

//分页参数
const page = ref({
    current: 1,
    pageSize: 10,
    total: 0,
    searchForm: {}
})
const tableRef = ref(null)
const list = ref([]);
const formRef = ref(null)
const formData = ref({
    "username": "",
    "nick_name": "",
    "email": "",
    "phone": "",
    "status": 0
})
const ids = ref([])
const formTitle = ref("")

//搜索列表
const searchList = ref([
    {
        label: '菜单名称',
        key: 'title',
        value: "",
        type: 'input',
        placeholder: '请输入菜单名称'
    },
    {
        label: '菜单跳转地址',
        key: 'path',
        value: "",
        type: 'input',
        placeholder: '请输入菜单跳转地址'
    },
    {
        label: '前端路径地址',
        key: 'component',
        value: "",
        type: 'input',
        placeholder: '请输入前端路径地址'
    },
    {
        label: '状态',
        key: 'status',
        value: "",
        type: 'select',
        placeholder: '请选择状态',
        options: [
            {
                label: '已禁用',
                value: 1
            },
            {
                label: '已启用',
                value: 2
            }],
    }]);

//表格列
const columns = [
    {
        title: '标识',
        dataIndex: 'id',
        width: 110,
        sortable: {
            sortDirections: ['ascend', 'descend']
        }
    },
    {
        title: '菜单名称',
        dataIndex: 'title',
        tooltip: true,
        ellipsis: true,
        width: 150,
        slotName: 'title'
    },
    {
        title: '菜单跳转路径',
        dataIndex: 'path',
        width: 150,
        tooltip: true,
        ellipsis: true,
    },
    {
        title: '前端文件路径',
        dataIndex: 'component',
        tooltip: true,
        ellipsis: true,
    },
    {
        title: '状态',
        dataIndex: 'status',
        width: 90,
        tooltip: true,
        ellipsis: true,
        slotName: 'status'
    },
    {
        title: '父节点',
        dataIndex: 'parent_id',
        width: 80,
        tooltip: true,
        ellipsis: true,
    }, {
        title: '排序',
        dataIndex: 'sort',
        width: 70,
        tooltip: true,
        ellipsis: true,
    }, {
        title: '创建时间',
        dataIndex: 'created_at',
        tooltip: true,
        ellipsis: true,
        width: 180,
        slotName: 'created_at'
    }, {
        title: '操作',
        slotName: 'optional',
        width: 240,
    }
];

//表单验证
const formRules = reactive({
    parent_id: [
        {required: true, message: '请选择父级菜单', trigger: 'blur'},
    ],
    title: [
        {required: true, message: '请输入菜单名称', trigger: 'blur'},
    ],
    is_cache: [
        {required: true, message: '请选择是否缓存', trigger: 'blur',},
    ],
    status: [
        {required: true, message: '请选择状态', trigger: 'blur',},
    ],
    is_tab: [
        {required: true, message: '请选择标签栏', trigger: 'blur',},
    ],
})

//表单列表
const formList = ref([
    {
        label: '父级菜单',
        key: 'parent_id',
        value: "",
        type: 'tree-select',
        placeholder: '请选择父级菜单',
        options: [],
    },
    {
        label: '菜单名称',
        key: 'title',
        value: "",
        type: 'input',
        placeholder: '请输入菜单名称'
    },
    {
        label: '菜单跳转地址',
        key: 'path',
        value: "",
        type: 'input',
        placeholder: '请输入菜单跳转地址'
    },
    {
        label: '前端路径地址',
        key: 'component',
        value: "",
        type: 'input',
        placeholder: '请输入前端跳转地址'
    },
    {
        label: '前端参数',
        key: 'param',
        value: "",
        type: 'input',
        placeholder: '请输入前端参数'
    },
    {
        label: '排序',
        key: 'sort',
        value: "",
        type: 'number',
        placeholder: '请输入排序'
    },
    {
        label: '菜单图标',
        key: 'icon',
        value: "",
        type: 'slot',
        placeholder: '请选择图标',
        options: [],
    },
    {
        label: '缓存页面',
        key: 'is_cache',
        value: "",
        type: 'radio-group-btn',
        placeholder: '请选择是否缓存',
        options: [
            {
                label: '关闭缓存',
                value: 1
            },
            {
                label: '启用缓存',
                value: 2
            }],
    },
    {
        label: '显示便签栏',
        key: 'is_tab',
        value: "",
        type: 'radio-group-btn',
        placeholder: '请选择是否缓存',
        options: [
            {
                label: '关闭便签栏',
                value: 1
            },
            {
                label: '开启便签栏',
                value: 2
            }],
    },
    {
        label: '状态',
        key: 'status',
        value: "",
        type: 'radio-group-btn',
        placeholder: '请选择状态',
        options: [
            {
                label: '禁用中',
                value: 1
            },
            {
                label: '启用中',
                value: 2
            }],
    }
]);

/**
 * 表单列表索引映射
 * @type {{}}
 */
const formListIndexMap = formList.value.reduce((map, item, index) => {
    map[item.key] = index;
    return map;
}, {});

/**
 * 搜索框列表索引映射
 * @type {{}}
 */
const searchListIndexMap = searchList.value.reduce((map, item, index) => {
    map[item.key] = index;
    return map;
}, {});

export {
    searchList,
    columns,
    page,
    formList,
    list,
    formRef,
    formData,
    formRules,
    ids,
    tableRef,
    formTitle,
    formListIndexMap,
    searchListIndexMap
};