// UMD 魔法代码
// if the module has no dependencies, the above pattern can be simplified to
(function (root, factory) {
    if (typeof define === 'function' && define.amd) {
        // AMD. Register as an anonymous module.
        define([], factory);
    } else if (typeof module === 'object' && module.exports) {
        // Node. Does not work with strict CommonJS, but
        // only CommonJS-like environments that support module.exports,
        // like Node.
        module.exports = factory();
    } else {
        // Browser globals (root is window)
        root.handle = factory();
    }
}(this, function () {
    /**
     * 初始化需要的leaflet
     * @param div_id
     * @param tianditu_key
     * @param center
     * @returns {leaflet_map}
     */
    var initLeafletMap = function initLeafletMap(div_id, tianditu_key, center=[39.928, 116.404]) {
        let TianDiTuNormalM = L.tileLayer.chinaProvider('TianDiTu.Normal.Map', {
                key: tianditu_key,
                maxZoom: 18, minZoom: 2
            }),
            TianDiTuNormalA = L.tileLayer.chinaProvider('TianDiTu.Normal.Annotion', {
                key: tianditu_key,
                maxZoom: 18, minZoom: 2
            }),
            TianDiTuImgM = L.tileLayer.chinaProvider('TianDiTu.Satellite.Map', {
                key: tianditu_key,
                maxZoom: 18, minZoom: 2
            }),
            TianDiTuImgA = L.tileLayer.chinaProvider('TianDiTu.Satellite.Annotion', {
                key: tianditu_key,
                maxZoom: 18, minZoom: 2
            }),
            BaiduNormalMap = L.tileLayer.chinaProvider('Baidu.Normal.Map', {
                maxZoom: 18, minZoom: 2
            }),
            BaiduSatelliteMap = L.tileLayer.chinaProvider('Baidu.Satellite.Map', {
                maxZoom: 18, minZoom: 2
            }),
            BaiduAnnotionMap = L.tileLayer.chinaProvider('Baidu.Satellite.Annotion', {
                maxZoom: 18, minZoom: 2
            }),
            GaoDeNormalM = L.tileLayer.chinaProvider('GaoDe.Normal.Map', {
                maxZoom: 18, minZoom: 2
            }),
            GaoDeImgM = L.tileLayer.chinaProvider('GaoDe.Satellite.Map', {
                maxZoom: 18, minZoom: 2
            }),
            GaoDeImgA = L.tileLayer.chinaProvider('GaoDe.Satellite.Annotion', {
                maxZoom: 18, minZoom: 2
            }),
            OSMNormalMap = L.tileLayer.chinaProvider('OSM.Normal.Map', {
                maxZoom: 18, minZoom: 2,
            }),
            GeoQNormal = L.tileLayer.chinaProvider('Geoq.Normal.Map', {
                maxZoom: 16, minZoom: 5
            }),
            GeoQPurplishBlueM = L.tileLayer.chinaProvider('Geoq.Normal.PurplishBlue', {
                maxZoom: 16, minZoom: 5
            }),
            GeoQGrayM = L.tileLayer.chinaProvider('Geoq.Normal.Gray', {
                maxZoom: 16, minZoom: 5
            }),
            GeoQWarmM = L.tileLayer.chinaProvider('Geoq.Normal.Warm', {
                maxZoom: 16, minZoom: 5
            }),
            GoogleNormalMap = L.tileLayer.chinaProvider('Google.Normal.Map', {
                maxZoom: 18,
                minZoom: 5
            }),
            GoogleSatelliteMap = L.tileLayer.chinaProvider('Google.Satellite.Map', {
                maxZoom: 18,
                minZoom: 5
            }),
            GoogleRouteMap = L.tileLayer.chinaProvider('Google.Satellite.Annotion', {
                maxZoom: 18,
                minZoom: 5
            }),
            MapBox = new L.TileLayer(
                '//api.mapbox.com/styles/v1/mapbox/streets-v10/tiles/{z}/{x}/{y}@2x?access_token=' +
                'sk.eyJ1Ijoic29uZ3RpYW5sdW4iLCJhIjoiY2w0Mms3aDduMDVuYTNmbWwxcnk3ZTFhZCJ9.0J3s1k_txufnnMpLqRf1UQ'
                , {
                attribution:
                    'Map data © <a href="http://osm.org/copyright">OpenStreetMap</a> contributors. Tiles from <a href="https://www.mapbox.com">Mapbox</a>.',
                tileSize: 512,
                zoomOffset: -1
            });

        let GaoDeNormal = L.layerGroup([GaoDeNormalM]),
            GaoDeImage = L.layerGroup([GaoDeImgM, GaoDeImgA]),
            BaiduImg = L.layerGroup([BaiduSatelliteMap, BaiduAnnotionMap]),
            TianDiNormal = L.layerGroup([TianDiTuNormalM, TianDiTuNormalA]),
            TianDiImage = L.layerGroup([TianDiTuImgM, TianDiTuImgA]);

        let baseLayers = {
            "天地图": TianDiNormal,
            "天地卫星图": TianDiImage,
            "高德地图": GaoDeNormal,
            "高德影像": GaoDeImage,
            "百度地图": BaiduNormalMap,
            "百度影像": BaiduSatelliteMap,
            "百度路线": BaiduImg,
            "OSM": OSMNormalMap,
            "GeoQ": GeoQNormal,
            "GeoQ午夜蓝": GeoQPurplishBlueM,
            "GeoQ灰色": GeoQGrayM,
            "GeoQ暖色": GeoQWarmM,
            "Google": GoogleNormalMap,
            "Google卫星": GoogleSatelliteMap,
            "Google路线": GoogleRouteMap,
            "MapBox": MapBox
        }

        let overlayLayers = {
            // "标注": BaiduAnnotionMap
        }
        // let crs = L.CRS.Baidu
        let crs = L.CRS.EPSG3857
        if (getMapName() === '百度地图' ||
            getMapName() === '百度影像' ||
            getMapName() === '百度路线') {
            crs = L.CRS.Baidu
        }

        let leaflet_map = L.map(div_id, {
            attributionControl: true,
            crs: crs,
            fullscreenControl: true,
            center: center,
            zoom: 13,
            layers: [baseLayers[getMapName()]],
            zoomControl: false,
            detectRetina: true
        });

        leaflet_map.attributionControl.setPrefix(
            "<a href='https://www.frytea.com'>🌱 Frytea</a> | " +
            "<a href='http://minepin.frytea.com'>📌 MinePin</a>");

        L.control.layers(baseLayers, overlayLayers).addTo(leaflet_map);
        L.control.zoom({
            zoomInTitle: '放大',
            zoomOutTitle: '缩小'
        }).addTo(leaflet_map);
        L.control.scale().addTo(leaflet_map);
        L.control.locate().addTo(leaflet_map);

        // 监听 layer 切换事件从而切换 crs
        leaflet_map.on("baselayerchange", function(e) {
            // console.log(e.name);
            setMapName(e.name);
            if (e.name === "百度地图" || e.name === "百度影像") {
                changeCRS(leaflet_map, L.CRS.Baidu)
            } else {
                changeCRS(leaflet_map, L.CRS.EPSG3857)
            }
        });

        return leaflet_map
    }

    let setMapName = function(map_name) {
        localStorage.setItem('minepin_map_name',map_name);
    }

    let getMapName = function() {
        if (localStorage.getItem('minepin_map_name')) {
            return localStorage.getItem('minepin_map_name');
        } else {
            return "MapBox";
        }
    }

    var getMaxBounds = function getMaxBounds(crs) {
        const { bounds } = crs.projection;
        return new L.LatLngBounds(
            crs.unproject(bounds.min),
            crs.unproject(bounds.max),
        );
    }

    var changeCRS = function changeCRS(map, crs) {
        const bounds = map.getBounds();
        map.options.crs = crs;
        // Ensure zoom is not affected by differing CRS scales
        map.options.zoomSnap = 0;
        map.fitBounds(bounds);
        map.setMaxBounds(crs instanceof L.Proj.CRS ? getMaxBounds(crs) : null);
        map.options.zoomSnap = 1;
    }

    var setLeafletCenter = function setLeafletCenter(lat,lng, map) {
        map.setView(new L.LatLng(lat, lng));
    }

    var addMarkerToLeaflet = function addMarkerToLeaflet(lat,lng, map, setCenter=false) {
        if (setCenter === true) {
            setLeafletCenter(lat,lng, map)
        }
        return L.marker([lat, lng]).addTo(map);
    }

    var removeMarkerLayer = function removeMarkerLayer(markerLayer, map) {
        if (markerLayer) {
            map.removeLayer(markerLayer);
        }
    }

    return {
        initLeafletMap: initLeafletMap,
        getMaxBounds: getMaxBounds,
        changeCRS: changeCRS,
        setLeafletCenter: setLeafletCenter,
        addMarkerToLeaflet: addMarkerToLeaflet,
        removeMarkerLayer: removeMarkerLayer
    }
}));