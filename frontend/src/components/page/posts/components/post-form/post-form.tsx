import { Button, Text, Textarea } from "~/components/ui";
import styles from "./post-form.module.scss";

type Props = {
	onChange: (e: React.ChangeEvent<HTMLTextAreaElement>) => void;
	onSubmit: () => void;
	disableSubmit?: boolean;
};

export const PostForm = ({ onChange, onSubmit }: Props) => {
	const handleChange = (e: React.ChangeEvent<HTMLTextAreaElement>) => {
		onChange(e);
	};
	const handleSubmit = () => {
		onSubmit();
	};

	const onKeyDown = (e: React.KeyboardEvent<HTMLTextAreaElement>) => {
		if ((e.key === "Enter" && e.ctrlKey) || e.metaKey) {
			onSubmit();
		}
	};

	return (
		<div>
			<div className={styles.form__input}>
				<Textarea rows={2} onChange={handleChange} onKeyDown={onKeyDown} />
				<Button variant="fill" color="primary" size="md" onClick={handleSubmit}>
					Post
				</Button>
			</div>
			<Text as="span" size="sm">
				Cmd + Enter で送信
			</Text>
		</div>
	);
};
