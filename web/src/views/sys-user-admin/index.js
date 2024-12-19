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
        label: '用户名',
        key: 'username',
        value: "",
        type: 'input',
        placeholder: '请输入姓名'
    },
    {
        label: '昵称',
        key: 'nick_name',
        value: "",
        type: 'input',
        placeholder: '请输入昵称'
    },
    {
        label: '电子邮箱',
        key: 'email',
        value: "",
        type: 'input',
        placeholder: '请输入电子邮箱'
    },
    {
        label: '手机号',
        key: 'phone',
        value: "",
        type: 'input',
        placeholder: '请输入手机号'
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
    },
    {
        label: '创建时间',
        key: 'created_at',
        value: [],
        type: 'dateRange',
        placeholder: '请选择创建时间',
    }]);

//表格列
const columns = [
    {
        title: '标识',
        dataIndex: 'id',
        width: 120,
        sortable: {
            sortDirections: ['ascend', 'descend']
        },
        visible: true
    },
    {
        title: '用户名',
        dataIndex: 'username',
        tooltip: true,
        ellipsis: true,
        visible: true
    },
    {
        title: '昵称',
        dataIndex: 'nick_name',
        tooltip: true,
        ellipsis: true,
        visible: true
    },
    {
        title: '手机号',
        dataIndex: 'phone',
        minWidth: 100,
        tooltip: true,
        ellipsis: true,
        visible: true
    },
    {
        title: '电子邮箱',
        dataIndex: 'email',
        minWidth: 100,
        tooltip: true,
        ellipsis: true,
        visible: true
    }, {
        title: '状态',
        dataIndex: 'status',
        width: 90,
        tooltip: true,
        ellipsis: true,
        visible: true,
        slotName: 'status'
    }, {
        title: '创建时间',
        dataIndex: 'created_at',
        tooltip: true,
        ellipsis: true,
        visible: true,
        minWidth: 100,
        slotName: 'created_at'
    }, {
        title: '操作',
        slotName: 'optional',
        width: 220,
        visible: true
    }
];

//表单验证
const formRules = reactive({
    username: [
        {required: true, message: '请输入用户名', trigger: 'blur'},
    ],
    nick_name: [
        {required: true, message: '请输入昵称', trigger: 'blur',},
    ],
    password: [
        {required: true, message: '请输入密码', trigger: 'blur',},
    ]
})

//表单列表
const formList = ref([
    {
        label: '用户名',
        key: 'username',
        value: "",
        type: 'input',
        placeholder: '请输入姓名'
    },
    {
        label: '密码',
        key: 'password',
        value: "",
        type: 'hidden',
        placeholder: '请输入密码'
    },
    {
        label: '昵称',
        key: 'nick_name',
        value: "",
        type: 'input',
        placeholder: '请输入昵称'
    },
    {
        label: '电子邮箱',
        key: 'email',
        value: "",
        type: 'input',
        placeholder: '请输入电子邮箱'
    },
    {
        label: '手机号',
        key: 'phone',
        value: "",
        type: 'input',
        placeholder: '请输入手机号'
    },
    {
        label: '状态',
        key: 'status',
        value: "",
        type: 'radio-group-btn',
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