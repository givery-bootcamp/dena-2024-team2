import { Button, Card } from "~/components/ui";
import styles from "./login.module.css";

export const LoginPage = () => {
	return (
		<div className={styles.root}>
			<Card
				header={<h1>Login</h1>}
				body={
					<div className={styles.card__body}>
						<input type="text" placeholder="Username" />
						<input type="password" placeholder="Password" />
						<Button size="md" color="primary" variant="fill">
							Login
						</Button>
					</div>
				}
			/>
		</div>
	);
};
