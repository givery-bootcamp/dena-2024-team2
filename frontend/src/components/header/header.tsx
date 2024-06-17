import { Text } from "../text";
import styles from "./header.module.scss";

export const Header = () => {
	return (
		<header className={styles.header}>
			<Text as="h1" size="xl" bold>
				team-2 app
			</Text>
		</header>
	);
};
