import styles from "./App.module.scss";
import { Header } from "./components/ui";

function App() {
	return (
		<div className={styles.root}>
			<Header />
		</div>
	);
}

export default App;
