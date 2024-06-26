import { Button, Card, Input } from "~/components/ui";
import styles from "./login.module.scss";

export const LoginPage = () => {
	const handleClickLogin = () => {
		console.log("Login");
	};

	return (
		<div className={styles.root}>
			<div className={styles.card}>
				<Card
					header={<h1>Login</h1>}
					body={
						<div className={styles.card__body}>
							<Input type="text" placeholder="Username" />
							<Input type="password" placeholder="Password" />
							<Button
								size="md"
								color="primary"
								variant="fill"
								onClick={handleClickLogin}
							>
								Login
							</Button>
						</div>
					}
				/>
			</div>
		</div>
	);
};
