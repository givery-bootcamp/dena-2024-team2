import { Header } from "./components/header";
import "modern-css-reset";
import styles from "./App.module.scss";

function App() {
	return (
		<div className={styles.root}>
			<Header />
		</div>
	);
}

export default App;
