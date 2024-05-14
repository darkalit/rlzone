import React, { useEffect, useState } from "react";
import { Link, useParams } from "react-router-dom";
import { GetItemByID } from "../../Services/Items";
import "./ViewItem.style.css";
import { ObjectTable } from "../../UI";
import Note from "../../Images/note.svg";
import Star from "../../Images/rateStar.svg";
import Credits from "../../Images/Credits.svg";

export default function ViewItem() {
  const { id } = useParams();

  const [item, setItem] = useState({});
  const [image, setImage] = useState("");
  const [stock, setStock] = useState({});
  useEffect(() => {
    (async function () {
      const data = await GetItemByID(id);
      const { Image, ID, Stock, ...rest } = data;
      setStock(Stock);
      setItem(rest);
      setImage(Image);
    })();
  }, [id]);

  // useEffect(() => {
  //   let itemCopy = { ...item };
  //   let image = itemCopy?.Image?.slice();
  //   setImage(image);
  //   delete itemCopy.Image;
  // }, [item]);

  const sidebarStyle = {
    width: "340px",
    padding: "0px 36px 0px 36px",
  };

  const creditsStyle = {
    width: "24px",
    height: "24px",
  };

  const carImg = {
    width: "254px",
  };

  return (
    <div className="view-item-page">
      <main className="main">
        <div className="main-content">
          <div className="main-content-sidebar" style={sidebarStyle}>
            <div className="card grey-container">
              {image && <img src={image} style={carImg} alt="car" />}
              <h1 className="car-name">{item.Name}</h1>
              {!stock?.Price ? (
                <p className="stock-item">Not in stock</p>
              ) : (
                <div className="credits">
                  <img src={Credits} alt="credits" style={creditsStyle} />
                  <p className="credit-amount">{stock?.Price}</p>
                </div>
              )}
              <div className="switch-block">
                <div className="alert-block">
                  <Link to="#" className="alert-text">
                    <p>Alert me</p>
                  </Link>
                </div>
                <div className="star-block">
                  <img className="star-rate" src={Star} alt="" />
                </div>
              </div>
            </div>
          </div>
          <div className=" main-content-block">
            <div className="main-description-table grey-container">
              <ObjectTable className="main-table" data={item} />
            </div>

            <div className="main-content-block-pages">
              <div className="main-block-description grey-container">
                <div className="block-description-title">
                  <img src={Note} alt="note" />
                  <h1 className="accent-text">Description</h1>
                </div>
                <p className="description-text">{stock?.Description}</p>
              </div>
            </div>

            <div className="pages-filler"></div>
          </div>
        </div>
      </main>
    </div>
  );
}
