import styles from "./input.module.scss";

type Props = React.InputHTMLAttributes<HTMLInputElement>;

export const Input = ({ name, value, onChange, ...rest }: Props) => {
	return (
		<input
			type="text"
			className={styles.input}
			name={name}
			id={name}
			value={value}
			onChange={onChange}
			{...rest}
		/>
	);
};
