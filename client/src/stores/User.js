import { defineStore } from "pinia";
import api from "@/plugins/axios";

export const useUserStore = defineStore("user-store", {
  state: () => ({
    users: [],
    selectedUser: null,
    error: "",
  }),
  getters: {
    getSelectedUser: (state) => state.selectedUser,
  },
  actions: {
    //setters
    async setSelectedUser(user) {
      this.selectedUser = user;
    },
    async getUsers() {
      try {
        const response = await api.get("/admin/user");
        this.users = response.data.data;
      } catch (error) {
        switch (error.response?.status) {
          case 400:
            this.error = "Invalid request. Please try again.";
            break;
          case 401:
            this.error = "Unauthorized. Please login.";
            break;
          case 404:
            this.error = "Company not found. Please try again.";
            break;
          default:
            this.error =
              "Something went wrong. Please try again or contact support.";
            break;
        }
        throw error;
      }
    },
  },
});
