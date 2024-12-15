import {fas} from "@fortawesome/free-solid-svg-icons";

const iconAll = Array.from(new Set(
    Object.keys(fas).map(iconKey => `fa-${fas[iconKey].iconName}`)
));

export default iconAll;