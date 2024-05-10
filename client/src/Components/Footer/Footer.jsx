import React from "react";
import "./Footer.style.css";
import { Link } from "react-router-dom";

export default function Footer() {
  return (
    <footer className="footer">
      <div className="footer-columns">
        <div className="footer-first-column">
          <img src="resources/img/footLogo.svg" alt="logo" />
          <p className="footer-number">+ 380(67) 246-28-32</p>
          <p className="footer-contact-email">support@rlzone.agency</p>
        </div>
        <div className="footer-second-column-block">
          <h1 className="footer-submain-text">Quick Links</h1>
          <div className="footer-block-links">
            <Link to="/about" className="footer-link">
              About us
            </Link>
            <Link to="/license" className="footer-link">
              User license
            </Link>
            <Link to="/privacy" className="footer-link">
              Privacy policy
            </Link>
            <Link to="/terms" className="footer-link">
              Terms of use
            </Link>
          </div>
        </div>
        <div className="footer-third-column">
          <h1 className="footer-submain-text">Get Notifications</h1>
          <div className="footer-block-subscribe">
            <input
              className="footer-subscribe"
              type="email"
              placeholder="Email"
            />
            <button type="button" className="footer-subscribe-button">
              ➜
            </button>
          </div>
        </div>
      </div>
      <div className="footer-separator"></div>
      <div className="footer-bottom-block">
        <div className="footer-socials">
          <img src="resources/img/Linkedin.svg" alt="Linkedin" />
          <img src="resources/img/facebook.svg" alt="facebook" />
          <img src="resources/img/Twitter.svg" alt="twitter" />
        </div>
        <div className="footer-product-logo">
          <p className="footer-text">A product of</p>
          <img src="resources/img/LogoFooter.svg" alt="LogoFooter" />
        </div>
        <div className="footer-copyright">
          <p className="footer-text">© 2024 RL Zone. All rights reserved</p>
        </div>
      </div>
    </footer>
  );
}
