// import { tw } from "@/utils/tailwind";

export function Frame({ children, className }: React.ComponentProps<"div">) {
	return (
		<div className={["mx-auto w-full max-w-2xl px-4", className].join(" ")}>
			{children}
		</div>
	);
}
