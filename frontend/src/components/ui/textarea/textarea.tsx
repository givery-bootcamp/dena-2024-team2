import styles from "./textarea.module.scss";

type Props = {
	rows?: number;
	defaultValue?: string;
	value?: string;
	placeholder?: string;
	onChange?: (e: React.ChangeEvent<HTMLTextAreaElement>) => void;
	onKeyDown?: (e: React.KeyboardEvent<HTMLTextAreaElement>) => void;
};

export const Textarea = ({
	rows = 3,
	defaultValue,
	value,
	placeholder,
	onChange,
	onKeyDown,
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
			onKeyDown={onKeyDown}
			placeholder={placeholder}
		/>
	);
};
