import styles from "./textarea.module.scss";

type Props = {
	rows?: number;
	defaultValue?: string;
	value?: string;
	placeholder?: string;
	onChange?: (e: React.ChangeEvent<HTMLTextAreaElement>) => void;
};

export const Textarea = ({
	rows = 3,
	defaultValue,
	value,
	placeholder,
	onChange,
}: Props) => {
	const handleChange = (e: React.ChangeEvent<HTMLTextAreaElement>) => {
		onChange?.(e);
	};
	return (
		<textarea
			className={styles.textarea}
			rows={rows}
			defaultValue={defaultValue}
			value={value}
			onChange={handleChange}
			placeholder={placeholder}
		/>
	);
};
