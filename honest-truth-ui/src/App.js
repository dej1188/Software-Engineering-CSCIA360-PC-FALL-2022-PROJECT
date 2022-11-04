import './App.css';

import { useSelector } from "react-redux";

import Video from "./components/video";
import Notations from "./components/notations";

function App() {

    const notations = useSelector((state) => state.notations.notations)

  return (
    <div className="App">
        <header>The Honest Truth</header>

        <div id="main-wrapper">

            <Video video={{id: "hwYzrSF9unk"}} />

            <Notations notations={notations} />

        </div>

    </div>
  );
}

export default App;
