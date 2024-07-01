import styles from "./card.module.scss";

type Props = {
	header: React.ReactNode;
	body: React.ReactNode;
	bodyRef?: React.RefObject<HTMLDivElement>;
};

export const Card = ({ header, body, bodyRef }: Props) => {
	return (
		<div className={styles.root}>
			<div className={styles.header}>{header}</div>
			<div className={styles.body} ref={bodyRef}>
				{body}
			</div>
		</div>
	);
};
