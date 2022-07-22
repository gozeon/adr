import axios from "axios";
import { defineStore } from "pinia";

export const useProjectStore = defineStore({
  id: "project",
  state: () => ({
    projects: [],
    loading: false,
    activeProject: {},
  }),
  getters: {},
  actions: {
    async loadProject() {
      this.loading = true;
      const res = await axios.get("/api/projects");
      this.projects = res.data?.data;
      this.loading = false;
    },
    async createOrUpdateProject(data: any) {
      this.loading = true;
      const res = await axios.post("/api/project", data);
      this.loading = false;
      return res;
    },
    async loadDetail(id: any) {
      this.loading = true;
      const res = await axios.get(`/api/project/${id}`);
      this.activeProject = res.data?.data;
      this.loading = false;
      return res;
    },
  },
});
