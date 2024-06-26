import styles from "./input.module.scss";

type Props = React.InputHTMLAttributes<HTMLInputElement>;

export const Input = ({ name, value, onChange, ...rest }: Props) => {
	const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
		onChange?.(e);
	};
	return (
		<input
			type="text"
			className={styles.input}
			name={name}
			id={name}
			value={value}
			onChange={handleChange}
			{...rest}
		/>
	);
};
