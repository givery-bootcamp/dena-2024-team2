import { createPortal } from "react-dom";
import { Button, Card, Input, Text } from "~/components/ui";
import styles from "./new-server-dialog.module.scss";

type Props = {
	open: boolean;
	onClose: () => void;
	onSubmit: () => void;
	onChangeName: (e: React.ChangeEvent<HTMLInputElement>) => void;
};

export const NewServerDialog = ({
	open,
	onClose,
	onSubmit,
	onChangeName,
}: Props) => {
	const handleClick = () => {
		onSubmit();
	};
	const handleChangeName = (e: React.ChangeEvent<HTMLInputElement>) => {
		onChangeName(e);
	};
	const handleClose = () => {
		onClose();
	};

	return (
		<>
			{open &&
				createPortal(
					<div className={styles.root}>
						<div className={styles.dialog}>
							<Card
								header={
									<Text as="h2" size="md" bold>
										Create a new server
									</Text>
								}
								body={
									<form>
										<Input
											type="text"
											onChange={handleChangeName}
											placeholder="Server Name"
											required
										/>
										<div className={styles.actions}>
											<Button
												type="submit"
												color="primary"
												variant="fill"
												size="md"
												onClick={handleClick}
											>
												Create
											</Button>
											<Button
												type="button"
												color="primary"
												variant="outline"
												size="md"
												onClick={handleClose}
											>
												Close
											</Button>
										</div>
									</form>
								}
							/>
						</div>
					</div>,
					document.querySelector("#root") as HTMLElement,
				)}
		</>
	);
};
