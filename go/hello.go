CREATE TABLE `z_docs_content_6` (
  `docid` varchar(32) NOT NULL DEFAULT '' COMMENT '文书ID',
  `trialdate` date NOT NULL DEFAULT '0000-00-00',
  `docno` varchar(32) NOT NULL DEFAULT '' COMMENT '案号',
  `city` varchar(16) NOT NULL DEFAULT '' COMMENT '城市',
  `psy_str` text NOT NULL,
  `courtlevel` varchar(16) NOT NULL DEFAULT '' COMMENT '法院基本名',
  `casetypeno` tinyint(2) NOT NULL DEFAULT '0' COMMENT '案件性质, 刑事 民事 ID',
  `courtsearchmsg` text NOT NULL,
  `wenshuendmsg` text NOT NULL,
  `fullcourtmsg` text NOT NULL,
  `c_cityid` varchar(2) NOT NULL DEFAULT '' COMMENT '城市ID',
  `accusermsg` text NOT NULL,
  `doctypeno` tinyint(2) NOT NULL DEFAULT '0' COMMENT '文书类型id',
  `trialmsg` text NOT NULL,
  `doctype` varchar(16) NOT NULL DEFAULT '' COMMENT '文书类型',
  `trialround` varchar(32) NOT NULL DEFAULT '' COMMENT '审理程序',
  `appelleemsg` text NOT NULL,
  `fuwenmsg` text NOT NULL,
  `judgmentmsg` text NOT NULL,
  `courtdeemmsg` text NOT NULL,
  `county` varchar(16) NOT NULL DEFAULT '' COMMENT '区县名',
  `sjy` text NOT NULL,
  `c_countyid` varchar(2) NOT NULL DEFAULT '' COMMENT '区县ID',
  `appellorstr` text NOT NULL,
  `courtid` varchar(16) NOT NULL DEFAULT '' COMMENT '法院ID',
  `fg_str` text NOT NULL COMMENT '法官信息',
  `province` varchar(16) NOT NULL DEFAULT '' COMMENT '省名',
  `courtname` varchar(32) NOT NULL DEFAULT '' COMMENT '法院名',
  `appellor` text NOT NULL,
  `basemsg` text NOT NULL,
  `trialroundno` tinyint(1) NOT NULL DEFAULT '0' COMMENT '审理程序编号',
  `casetypestr` varchar(16) NOT NULL DEFAULT '' COMMENT '案件性质, 刑事 民事',
  `c_provinceid` varchar(2) NOT NULL DEFAULT '' COMMENT '省份ID',
  PRIMARY KEY (`docid`),
  KEY `courtid` (`courtid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

api

    public function listJudge(){
        $p = p('p',1,'intval');
        if($p < 1){
            $p = 1;
        }
        $fyid = p('fyid');
        $fgid = p('fgid');
        $listRe =  $this->cache(300)->listJudge($p,10,$fyid,$fgid);
        $listRe = is_array($listRe)?revalue($listRe['code'],'',$listRe['data']):$listRe;
        return $this->result($listRe->getData(),$listRe->getCode(),$listRe->getMsg());

    }