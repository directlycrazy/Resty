import { Button } from "@/components/Button"
import { Card } from "@/components/Card"
import { Input } from "@/components/Input"

export default function Login() {
	return (
		<>
			<div className="flex min-h-svh flex-col items-center justify-center bg-muted p-6 md:p-10">
				<div className="w-full max-w-sm md:max-w-3xl">
					<div className="flex flex-col gap-6">
						<Card className="overflow-hidden !p-0 no-padding">
							<div className="grid p-0 md:grid-cols-2">
								<form className="p-6 md:p-8 md:py-10">
									<div className="flex flex-col gap-2">
										<div className="flex flex-col items-center text-center mb-3">
											<h1 className="text-2xl font-bold tracking-tight">Welcome back</h1>
											<p className="text-balance text-muted-foreground">Login to your Resty dashboard</p>
										</div>
										<div className="grid gap-2">
											<label className="text-sm text-muted-foreground" htmlFor="email">Email</label>
											<Input
												id="email"
												type="email"
												autoComplete="off"
												placeholder="resty@example.com"
												required />
										</div>
										<div className="grid gap-2">
											<label className="text-sm text-muted-foreground" htmlFor="password">Password</label>
											<Input
												id="password"
												type="password"
												autoComplete="off"
												placeholder="••••••••"
												required />
										</div>
										<Button type="submit" className="w-full">Login</Button>
									</div>
								</form>
								<div className="relative hidden bg-muted md:block inset-0 h-full w-full object-cover brightness-75 dark:brightness-50 bg-[url(/img/login.webp)]"></div>
							</div>
						</Card>
						<div className="text-balance text-center text-xs text-muted-foreground [&_a]:underline [&_a]:underline-offset-4 hover:[&_a]:text-primary">github.com/directlycrazy/Resty</div>
					</div>
				</div>
			</div>
		</>
	)
}