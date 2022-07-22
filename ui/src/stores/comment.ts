import axios from "axios";
import { defineStore } from "pinia";

export const useCommentStore = defineStore({
  id: "comment",
  state: () => ({
    comments: [],
    loading: false,
    activeComments: [],
  }),
  getters: {},
  actions: {
    async loadCommentById(id: any) {
      this.loading = true;
      const res = await axios.get(`/api/comment/${id}`);
      this.activeComments = res.data?.data;
      this.loading = false;
    },
    async createOrUpdateComment(data: any) {
      this.loading = true;
      const res = await axios.post("/api/comment", data);
      this.loading = false;
      return res;
    },
  },
});
