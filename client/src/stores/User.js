import { defineStore } from "pinia";
import axios from "axios";
import api from "@/plugins/axios";

export const useUserStore = defineStore("user-store", {
  state: () => ({
    error: "",
    status: "",
    total: 0,
    item: "",
    items: [],
  }),
  getters: {
    getError: (state) => state.error,
    getStatus: (state) => state.status,
    getTotal: (state) => state.total,
    getItem: (state) => state.item,
    getItems: (state) => state.items,
  },
  actions: {
    async createCompany(payload, router) {
      this.error = "";
      const data = {
        companyName: payload.companyName,
        companySize: payload.companySize,
        adminEmail: payload.adminEmail,
        password: payload.password,
        confirmPassword: payload.confirmPassword,
      };
      try {
        const response = await api.post("/company", data);
        if (response.status >= 200 && response.status < 300) {
          return true;
        } else {
          return false;
        }
      } catch (error) {
        if (error.response) {
          if (error.response.status === 409) {
            this.error = "User already exists. Please login.";
            setTimeout(() => {
              router.push("/auth/signin");
            }, 3000);
          } else {
            this.error =
              "Something went wrong. Please try again or contact support.";
          }
        }
        throw error;
      }
    },
    async signin(payload) {
      this.error = "";

      let data = {
        email: payload.email,
        password: payload.password,
      };

      try {
        const response = await api.post("/auth/signin", data);
        if (response.status >= 200 && response.status < 300) {
          return true;
        } else {
          return false;
        }
      } catch (error) {
        try {
          const response = await axios.post(api_url + "/auth/signin", data);
          if (response.status >= 200 && response.status < 300) {
            return true;
          } else {
            return false;
          }
        } catch (error) {
          switch (error.response?.status) {
            case 401:
            case 404:
              this.error = "Invalid email or password. Please try again.";
              break;
            case 403:
              this.error = "Please verify your email address.";
              break;
            default:
              this.error =
                "Something went wrong. Please try again or contact support.";
              break;
          }
          throw error;
        }
      }
    },
    async getCsrfToken() {
      try {
        const response = await api.get("/auth/csrf");
        if (response.status >= 200 && response.status < 300) {
          return true;
        }
      } catch (err) {
        return err;
      }
    },
  },
});
