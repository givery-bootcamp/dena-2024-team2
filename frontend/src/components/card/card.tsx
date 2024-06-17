import styles from "./card.module.scss";

type Props = {
	header: React.ReactNode;
	body: React.ReactNode;
};

export const Card = ({ header, body }: Props) => {
	return (
		<div className={styles.root}>
			<div className={styles.header}>{header}</div>
			<div className={styles.body}>{body}</div>
		</div>
	);
};
