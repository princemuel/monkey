import { Form, Outlet } from "react-router";
import { Theme, useTheme } from "remix-themes";

import { IconMoon, IconSun } from "@/assets/icons";
import {
	ImageDarkLG,
	ImageDarkSM,
	ImageLightLG,
	ImageLightSM,
} from "@/assets/images";
import { Frame } from "@/components/frame";
import type { Route } from "./+types/layout";

export default function Layout({ loaderData: _ }: Route.ComponentProps) {
	const [theme, setTheme] = useTheme();
	return (
		<>
			<header
				className="flex flex-col gap-12 py-32 bg-no-repeat bg-(image:--image-sm) sm:bg-(image:--image-xl) bg-cover"
				style={
					{
						"--image-xl":
							theme === Theme.DARK
								? `url(${ImageDarkLG})`
								: `url(${ImageLightLG})`,
						"--image-sm":
							theme === Theme.DARK
								? `url(${ImageDarkSM})`
								: `url(${ImageLightSM})`,
					} as React.CSSProperties
				}
			>
				<Frame className="flex items-center justify-between">
					<h1 className="uppercase font-bold tracking tracking-widest text-4xl text-white">
						Taska
					</h1>

					<button
						type="button"
						className="text-white"
						onClick={() =>
							void setTheme((value) =>
								value === Theme.DARK ? Theme.LIGHT : Theme.DARK,
							)
						}
					>
						<span className="sr-only">Toggle theme</span>
						{theme === Theme.DARK ? <IconSun /> : <IconMoon />}
					</button>
				</Frame>

				<Form action="/events" method="post">
					<input type="text" />
				</Form>
			</header>

			<Outlet />

			<footer className="py-12">
				<Frame className="flex items-center justify-center">
					<p className="text-base text-gray-400">
						Drag and drop to reorder list
					</p>
				</Frame>
			</footer>
		</>
	);
}
