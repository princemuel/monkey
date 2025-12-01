import type { RouteConfig } from "@react-router/dev/routes";
import { index, layout, prefix, route } from "@react-router/dev/routes";

export default [
  layout("components/layout.tsx", [index("./routes/_index.tsx")]),
  ...prefix("api", [route("theme", "./routes/api/theme.ts")]),
] satisfies RouteConfig;
