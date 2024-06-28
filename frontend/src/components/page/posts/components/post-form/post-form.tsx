import { Button, Text, Textarea } from "~/components/ui";
import styles from "./post-form.module.scss";

type Props = {
	onChange: (e: React.ChangeEvent<HTMLTextAreaElement>) => void;
	onSubmit: () => void;
};

export const PostForm = ({ onChange, onSubmit }: Props) => {
	const handleChange = (e: React.ChangeEvent<HTMLTextAreaElement>) => {
		onChange(e);
	};
	const handleSubmit = () => {
		onSubmit();
	};

	return (
		<div>
			<div className={styles.form__input}>
				<Textarea rows={2} onChange={handleChange} />
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
