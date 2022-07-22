import { defineConfig } from "windicss/helpers";
import formPlugin from "windicss/plugin/forms";
export default defineConfig({
  darkMode: "class",
  theme: {
    extend: {
      colors: {
        // https://react-notion-x-demo.transitivebullsh.it/54bf56611797480c951e5c1f96cb06f2
        notion: {
          red: "#e03e3e",
          pink: "#ad1a72",
          blue: "#0b6e99",
          purple: "#6940a5",
          teal: "#4d6461",
          yellow: "#dfab01",
          orange: "#d9730d",
          brown: "#64473a",
          gray: "#9b9a97",
          red_background: "#fbe4e4",
          pink_background: "#f4dfeb",
          blue_background: "#ddebf1",
          purple_background: "#eae4f2",
          teal_background: "#ddedea",
          yellow_background: "#fbf3db",
          orange_background: "#faebdd",
          brown_background: "#e9e5e3",
          gray_background: "#ebeced",
        },
      },
    },
  },
  plugins: [formPlugin],
});
