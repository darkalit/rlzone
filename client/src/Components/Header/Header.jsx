import React from "react";
import "./Header.style.css";
import { useLocation, Link } from "react-router-dom";
import Button, { ButtonType } from "../../UI/Button/Button";

function UserView() {
  let path = useLocation().pathname;
  const selection = "header-section-selected";

  return (
    <>
      <div className="header-buy-sell">
        <div
          className={`header-section ${
            path.startsWith("/buy") ? selection : ""
          }`}
        >
          <Link to="/buy">Buy</Link>
        </div>
        <div
          className={`header-section ${
            path.startsWith("/inventory") ? selection : ""
          }`}
        >
          <Link to="/inventory">Inventory</Link>
        </div>
      </div>
      <div className="header-right-block">
        <div className="header-block-settings button-add-scale">
          <a className="header-svg-icons" href="basket.html">
            <img src="resources/img/CartBlock.svg" alt="Cart" />
          </a>
        </div>
        <div className="header-block-settings">
          <img
            src="resources/img/Credits.svg"
            alt="Credits"
            className="credits"
          />
          <p className="header-main-text">1290</p>
        </div>
        <div className="header-block-settings button-add-scale">
          <a className="header-svg-icons" href="profile.html">
            <img
              src="resources/img/profileIcon.svg"
              alt="profileIcon"
              className="profile-pic"
            />
          </a>
          <a href="profile.html">
            <p className="header-main-text">SampleName</p>
          </a>
        </div>
      </div>
    </>
  );
}

function AdminView() {
  let path = useLocation().pathname;
  const selection = "header-section-selected";

  return (
    <>
      <div className="header-buy-sell">
        <div
          className={`header-section ${
            path.startsWith("/listings") ? selection : ""
          }`}
        >
          <Link to="/listings">Edit listings</Link>
        </div>
        <div
          className={`header-section ${
            path.startsWith("/users") ? selection : ""
          }`}
        >
          <Link to="/users">User control</Link>
        </div>
      </div>
      <div className="header-right-block">
        <div className="header-block-settings button-add-scale">
          <Link to="/profile" className="header-svg-icons">
            <img
              src="resources/img/profileIcon.svg"
              alt="profileIcon"
              className="profile-pic"
            />
          </Link>
          <p className="header-main-text">AdminUser</p>
        </div>
      </div>
    </>
  );
}

function GuestView() {
  return (
    <div className="header-btns">
      <Button text="Login" type={ButtonType.Outline} width="106px" />
      <Button text="Sign Up" type={ButtonType.Fill} width="106px" />
    </div>
  );
}

export default function Header(props) {
  let view = null;

  switch (props.role) {
    case "User":
      view = UserView();
      break;
    case "Admin":
      view = AdminView();
      break;
    default:
      view = GuestView();
      break;
  }

  return (
    <header className="header">
      <div className="header-content">
        <a href="index.html">
          <img
            className="header-logo"
            src="resources/img/LogoMain.svg"
            alt="logo"
          />
        </a>
        <div className="header-block-search">
          <img src="resources/img/MagnifyingGlass.svg" alt="magnifyingGlass" />
          <input
            className="header-search"
            type="text"
            placeholder="Find an item..."
          />
          <button type="button" className="header-search-button">
            ➜
          </button>
        </div>
        {view}
      </div>
    </header>
  );
}
