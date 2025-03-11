import { RouteRecordRaw } from "vue-router"

//主页路由
export const HomeRoutes:RouteRecordRaw[] =[
  //主页
  { 
    path: '/',
    redirect: '/home',
    component: () => import("@/layout/HomeLayout.vue"),
    meta: { hidden:true },
    children:[
      {
        path: 'home',
        name: 'Home',
        component: () => import('@/views/home/Home.vue'),
        meta: { title: '首页', activeMenu: '/home',sort:1 }
      },
      {
        path: 'submitVuln',
        name: 'submitVuln',
        component: () => import('@/views/home/SubmitVuln.vue'),
        meta: { title: '提交漏洞', activeMenu: '/submitVuln',sort:1 }
      },
      {
        path: 'article',
        name: 'article',
        component: () => import('@/views/home/Article.vue'),
        meta: { title: '技术培训', activeMenu: '/article',sort:1 }
      },
      {
        path: 'articleDetail',
        name: 'ArticleDetail',
        component: () => import('@/views/home/ArticleDetail.vue'),
        meta: { title: '文章详情', activeMenu: '/home/ArticleDetail',icon:"carbon:blog" ,hidden:true}
      },
    ]
  },
]


//用户中心路由
export const UserCenterRoutes:RouteRecordRaw[] =[
  {
    path: '/user', 
    name: 'User',
    redirect: '/user/profile',
    component:  () => import("@/layout/UserCenterLayout.vue"),
    meta:{
        hidden:true,
        isTopLevel:true
    },
    children:[
       {
        path: 'profile',
        name: 'Profile',
          component: () => import('@/views/user/Profile.vue'),
          meta: { title: '我的信息', activeMenu: '/user/profile',icon:"tabler:user-filled" }
      },
    ]
  },
]

//厂商中心路由
export const MunaCenterRoutes:RouteRecordRaw[] =[

]


//管理员路由
export const AdminRoutes:RouteRecordRaw[] =[
    {
        path: '/admin', 
        name: 'Admin',
        redirect: '/admin/dashboard',
        component:  () => import("@/layout/AdminLayout.vue"),
        meta:{
            hidden:true,
            isTopLevel:true
        },
        children:[
           {
             path: 'dashboard',
             name: 'Dashboard',
               component: () => import('@/views/admin/Dashboard.vue'),
               meta: { title: '仪表盘', activeMenu: '/dashboard',icon:"material-symbols:home-rounded",sort:1 }
           },
           {
            path: 'profile',
            name: 'Profile',
              component: () => import('@/views/admin/Profile.vue'),
              meta: { title: '个人信息', activeMenu: '/admin/profile',hidden:true }
          },
           { 
             path: 'user',
             name: 'Userlist',
             component: () => import('@/views/admin/user/UserList.vue'),
               meta: { title: '用户管理', activeMenu: '/user/list',icon:"ix:user-management-filled",sort:2 }
           },
            { 
              path: 'system',
              name: 'System',
              meta: { title: '系统管理', activeMenu: '/system',icon:"material-symbols:settings-rounded",sort:3 },
              children:[
                  {
                      path: 'dictType',
                      name: 'dictType',
                      component: () => import('@/views/admin/system/DictType.vue'),
                      meta: { title: '数据字典', activeMenu: '/system/dictType',icon:"arcticons:thai-dict" }
                    },
                  {
                    path: 'sysConfig',
                    name: 'Vulnlist',
                    component: () => import('@/views/admin/system/SysConfig.vue'),
                    meta: { title: '系统设置', activeMenu: '/system/sysConfig',icon:"mynaui:config" }
                  },
                  {
                      path: 'dictData',
                      name: 'dictData',
                      component: () => import('@/views/admin/system/DictData.vue'),
                      meta: { title: '字典数据', activeMenu: '/system/dictData',icon:"arcticons:thai-dict",hidden:true}
                  },
              ]
            },
            { 
              path: 'article',
              name: 'Artilce',
              meta: { title: '文章管理', activeMenu: '/article',icon:"material-symbols:article",sort:2 },
              children:[
                {
                  path: 'articleList',
                  name: 'ArticleList',
                  component: () => import('@/views/admin/article/ArticleList.vue'),
                  meta: { title: '文章列表', activeMenu: '/article/ArticleList',icon:"ooui:articles-ltr" }
                },
                  {
                      path: 'publishArtcile',
                      name: 'PublishArticle',
                      component: () => import('@/views/admin/article/PublishArticle.vue'),
                      meta: { title: '发表文章', activeMenu: '/article/PublishArticle',icon:"dashicons:welcome-write-blog" }
                    },
                    {
                      path: 'articleGroup',
                      name: 'ArticleGroup',
                      component: () => import('@/views/admin/article/ArticleGroup.vue'),
                      meta: { title: '文章分栏', activeMenu: '/article/ArticleGroup',icon:"carbon:blog" }
                    },
                   
              ]
            },
            { 
              path: 'vuln',
              name: 'Vuln',
              meta: { title: '漏洞管理', activeMenu: '/article',icon:"ion:bug",sort:2 },
              children:[
                {
                  path: 'vulnType',
                  name: 'VulnType',
                  component: () => import('@/views/admin/vuln/VulnType.vue'),
                  meta: { title: '漏洞类型', activeMenu: '/vuln/VulnType',icon:"streamline:bug-virus-browser" }
                },
                {
                  path: 'updateVuln',
                  name: 'UpdateVuln',
                  component: () => import('@/views/admin/vuln/UpdateVuln.vue'),
                  meta: { title: '更新漏洞', activeMenu: '/vuln/updateVuln',icon:"lsicon:submit-filled",hidden:true }
                },
                {
                  path: 'vulnList',
                  name: 'VulnList',
                  component: () => import('@/views/admin/vuln/VulnList.vue'),
                  meta: { title: '漏洞列表', activeMenu: '/vuln/vulnList',icon:"gg:list" }
                },
                {
                  path: 'vulnDetail',
                  name: 'VulnDetail',
                  component: () => import('@/views/admin/vuln/VulnDetail.vue'),
                  meta: { title: '漏洞详情', activeMenu: '/vuln/vulnDetail',icon:"gg:list",hidden:true }
                },
              ]
            },
        ]
      },
]

