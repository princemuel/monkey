import { IconCheck, IconCross } from "@/assets/icons";
import { Frame } from "@/components/frame";
import { tasks } from "@/lib/database";
import { NavLink } from "react-router";
import type { Route } from "./+types/_index";

export function meta(_args: Route.MetaArgs) {
	return [
		{ title: "Taska" },
		{ name: "description", content: "Welcome to React Router!" },
	];
}

export async function loader(_args: Route.LoaderArgs) {
	return { tasks };
}

export default function Page({ loaderData }: Route.ComponentProps) {
	const { tasks } = loaderData;
	return (
		<main className="-my-10">
			<Frame className="flex flex-col gap-6">
				<section className="flex flex-col rounded-lg bg-white dark:bg-grey-900">
					<ul className="flex flex-col justify-around">
						{tasks.map((task) => (
							<li
								key={task.id}
								className="bg-transparent group border-b border-grey-100 dark:border-grey-800 flex items-center gap-4 px-8 py-6"
								data-checked={task.completed}
							>
								<button
									type="button"
									className="rounded-full p-2 bg-linear-120 bg-accent-600"
									data-checked={task.completed}
								>
									<span className="sr-only">toggle the task's status </span>
									<IconCheck className="in-data-[checked=true]:visible invisible text-white" />
								</button>
								<p
									data-checked={task.completed}
									className="flex-1 text-gray-600 dark:text-grey-200"
								>
									{task.text}
								</p>
								<button
									type="button"
									className="text-grey-400 dark:text-grey-800"
								>
									<span className="sr-only">delete task</span>
									<IconCross className="" />
								</button>
							</li>
						))}
					</ul>

					<div className="flex items-center justify-between px-8 py-6 gap-4">
						<p className="text-grey-400">
							{tasks.length} item{tasks.length === 1 ? "" : "s"} left
						</p>

						<section className="hidden md:block">
							<nav className="flex items-center gap-4 capitalize font-bold">
								{[
									{ href: "/", text: "All" },
									{ href: "/active", text: "Active" },
									{ href: "/completed", text: "Completed" },
								].map((route) => (
									<NavLink
										key={route.text}
										to={route.href}
										className={({ isActive }) =>
											isActive ? "text-accent-500" : "text-grey-400"
										}
									>
										{route.text}
									</NavLink>
								))}
							</nav>
						</section>

						<button type="button" className="text-grey-400">
							Clear Completed
						</button>
					</div>
				</section>

				<section className="md:hidden rounded-lg bg-white dark:bg-grey-900">
					<nav className="flex items-center gap-4 font-bold capitalize py-6 justify-center">
						{[
							{ href: "/", text: "all" },
							{ href: "/active", text: "active" },
							{ href: "/completed", text: "completed" },
						].map((route) => (
							<NavLink
								key={route.text}
								to={route.href}
								className={({ isActive }) =>
									isActive ? "text-accent-500" : "text-grey-400"
								}
							>
								{route.text}
							</NavLink>
						))}
					</nav>
				</section>
			</Frame>
		</main>
	);
}
