/** biome-ignore-all assist/source/organizeImports: just ignored */
import { themeSessionResolver } from "@/sessions.server";
import { createThemeAction } from "remix-themes";

export const action = createThemeAction(themeSessionResolver);
