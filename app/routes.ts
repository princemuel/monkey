import type { RouteConfig } from "@react-router/dev/routes";
import { index, layout, prefix, route } from "@react-router/dev/routes";

export default [
	layout("components/layout.tsx", [index("./routes/_index.tsx")]),
	...prefix("actions", [route("set-theme", "./routes/actions.set-theme.ts")]),
] satisfies RouteConfig;
