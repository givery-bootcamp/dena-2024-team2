import { Button, Card, Input, Text } from "~/components/ui";
import { useLoginForm } from "./hooks/use-login-form";
import styles from "./login.module.scss";

export const LoginPage = () => {
	const {
		handleChangePassword,
		handleChangeUsername,
		handleClickLogin,
		status,
	} = useLoginForm();
	return (
		<div className={styles.root}>
			<div className={styles.card}>
				<Card
					header={
						<Text as="h2" size="lg" bold>
							Login
						</Text>
					}
					body={
						<div className={styles.card__body}>
							<Input
								type="text"
								placeholder="Username"
								onChange={handleChangeUsername}
							/>
							<Input
								type="password"
								placeholder="Password"
								onChange={handleChangePassword}
							/>
							<Button
								size="md"
								color="primary"
								variant="fill"
								onClick={handleClickLogin}
								disabled={status === "pending"}
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
