import { IconMoon, IconSun } from "@/assets/icons";
import { Theme, useTheme } from "remix-themes";

export function ThemeSwitch() {
  const [theme, setTheme] = useTheme();
  return (
    <button
      type="button"
      className="text-white"
      onClick={() =>
        void setTheme((value) => (value === Theme.DARK ? Theme.LIGHT : Theme.DARK))
      }
    >
      <span className="sr-only">Toggle theme</span>
      {theme === Theme.DARK ? <IconSun /> : <IconMoon />}
    </button>
  );
}
