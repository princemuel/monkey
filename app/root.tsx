import {
	data,
	isRouteErrorResponse,
	Links,
	type LinksFunction,
	Meta,
	Outlet,
	Scripts,
	ScrollRestoration,
	useRouteLoaderData,
} from "react-router";
import {
	PreventFlashOnWrongTheme,
	Theme,
	ThemeProvider,
	useTheme,
} from "remix-themes";

import "@fontsource-variable/josefin-sans";

import type { Route } from "./+types/root";
import "./global.css";
import { themeSessionResolver } from "./sessions.server";

export const links: LinksFunction = () => {
	return [{ rel: "icon", type: "image/png", href: "/favicon.png" }];
};

export async function loader({ request }: Route.LoaderArgs) {
	const { setTheme, getTheme, commit } = await themeSessionResolver(request);

	const theme = getTheme() ?? Theme.LIGHT;
	setTheme(theme);

	return data(
		{ theme: getTheme() },
		{ headers: { "Set-Cookie": await commit() } },
	);
}

export function Layout({ children }: { children: React.ReactNode }) {
	const data = useRouteLoaderData("root") as Route.ComponentProps["loaderData"];
	return (
		<ThemeProvider specifiedTheme={data.theme} themeAction="/actions/set-theme">
			<Root ssrTheme={Boolean(data.theme)}> {children}</Root>
		</ThemeProvider>
	);
}

function Root({
	ssrTheme,
	children,
}: {
	ssrTheme: boolean;
	children: React.ReactNode;
}) {
	const [theme] = useTheme();
	return (
		<html lang="en" className={theme ?? ""}>
			<head>
				<meta charSet="utf-8" />
				<meta name="viewport" content="width=device-width, initial-scale=1" />
				<PreventFlashOnWrongTheme ssrTheme={ssrTheme} />
				<Meta />
				<Links />
			</head>
			<body className="relative grid min-h-svh grid-rows-[auto_1fr_auto] dark:bg-grey-950 bg-grey-50 w-full flex-col font-sans text-base md:text-lg">
				{children}
				<ScrollRestoration />
				<Scripts />
			</body>
		</html>
	);
}

export default function App() {
	return <Outlet />;
}

export function ErrorBoundary({ error }: Route.ErrorBoundaryProps) {
	let message = "Oops!";
	let details = "An unexpected error occurred.";
	let stack: string | undefined;

	if (isRouteErrorResponse(error)) {
		message = error.status === 404 ? "404" : "Error";
		details =
			error.status === 404
				? "The requested page could not be found."
				: error.statusText || details;
	} else if (import.meta.env.DEV && error && error instanceof Error) {
		details = error.message;
		stack = error.stack;
	}

	return (
		<main className="pt-16 p-4 container mx-auto">
			<h1>{message}</h1>
			<p>{details}</p>
			{stack && (
				<pre className="w-full p-4 overflow-x-auto">
					<code>{stack}</code>
				</pre>
			)}
		</main>
	);
}
