// function middleware(to, from, next) {
//   if (sessionStorage.getItem("accessToken")) {
//     next();
//   } else {
//     next("/");
//   }
// }

const originRoutes = [
  {
    isActive: true,
    path: "/",
    name: "DashboardLayout",
    component: () => import("../layouts/default.vue"),
    children: [
      {
        isActive: true,
        path: "",
        name: "Dashboard",
        component: () => import("../pages/index.vue"),
      },
    ],
  },
  {
    isActive: true,
    path: "/admin",
    name: "AdminLayout",
    component: () => import("../layouts/default.vue"),
    children: [
      {
        isActive: true,
        path: "",
        name: "AdminDashboard",
        component: () => import("../pages/admin/index.vue"),
      },
      {
        isActive: true,
        path: "company",
        name: "AdminCompany",
        component: () => import("../pages/admin/company.vue"),
      },
      {
        isActive: true,
        path: "locations",
        name: "AdminLocation",
        component: () => import("../pages/admin/location.vue"),
      },
      {
        isActive: true,
        path: "departments",
        name: "AdminDepartment",
        component: () => import("../pages/admin/department.vue"),
      },
      {
        isActive: true,
        path: "positions",
        name: "AdminPosition",
        component: () => import("../pages/admin/position.vue"),
      },
    ],
  },
  {
    isActive: false,
    path: "/onboarding",
    name: "OnboardingLayout",
    component: () => import("../layouts/full-screen.vue"),
    children: [
      {
        isActive: true,
        path: "",
        name: "Onboarding",
        component: () => import("../pages/onboarding/index.vue"),
      },
      {
        isActive: true,
        path: "confirmation",
        name: "Confirmation",
        component: () => import("../pages/onboarding/confirmation.vue"),
      },
    ],
  },
  {
    isActive: true,
    path: "/auth",
    name: "FullScreenLayout",
    component: () => import("@/layouts/full-screen.vue"),
    children: [
      {
        isActive: true,
        path: "signin",
        name: "SignIn",
        component: () => import("@/pages/auth/signin.vue"),
      },
      // {
      //   isActive: false,
      //   path: "signup",
      //   name: "SignUp",
      //   component: () => import("@/pages/auth/signup.vue"),
      // },
    ],
  },
  /* {
      isActive: true,
      path: '/:pathMatch(.*)*',
      name: 'NotFound',
      component: () => import('../pages/NotFound.vue')
} */
];

export const filterByActive = () => {
  const recuresiveMap = (routes = originRoutes) => {
    return routes
      .map((_route) => {
        if (_route.isActive) {
          return _route;
        }
      })
      .filter((v) => !!v);
  };

  return recuresiveMap();
};

export const routes = () => filterByActive();
