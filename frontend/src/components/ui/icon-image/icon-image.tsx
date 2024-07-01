import clsx from "clsx";
import placeHolder from "~/assets/place-holder.png";
import styles from "./icon-image.module.scss";

type Props = {
	src: string;
	size?: "sm" | "md" | "lg";
};

export const IconImage = ({ src, size = "md" }: Props) => {
	// urlへのアクセスが失敗したらplace-holderを表示する
	// 画像の読み込みに失敗した場合に発火するイベント
	const handleImageError = (
		e: React.SyntheticEvent<HTMLImageElement, Event>,
	) => {
		const img = e.target as HTMLImageElement;
		img.src = placeHolder;
	};

	return (
		<div className={clsx(styles.root, styles[size])}>
			<img
				className={styles.icon}
				src={src}
				alt="icon"
				onError={handleImageError}
			/>
		</div>
	);
};
