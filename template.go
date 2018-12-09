package main

import (
	"strings"
	"time"
)

func prepareEmailBody(params map[string]string) string {
	currentTime := time.Now().UTC()
	dateTime := currentTime.Format("02 Jan, 2006 15:04 pm")

	// Worst way of mergin the strings
	var formattedParams []string
	for k, v := range params {
		formattedParams = append(formattedParams, `
		<tr>
		<td width="120px" style="font-weight:bold;vertical-align: top;border-top: solid 1px #ededed;padding-top: 10px">`+k+`</td>
		<td style="border-top: solid 1px #ededed;padding-top: 10px">`+v+`</td>
		</tr>
		`)
	}

	damn := strings.Join(formattedParams, "")

	htmlBody := `
	<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
  <meta http-equiv="Content-Type" content="text/html; charset=3Dutf-8">
  <title>New form submission!</title>
</head>

<body yahoo="" bgcolor="#f9fafe" style="min-width: 100% !important; padding: 0; margin:0;">

  <table width="100%" bgcolor="#f9fafe" border="0" cellpadding="0" cellspacing="0">
    <tbody><tr>
      <td>
        
        <table bgcolor="#ffffff" class="m_-7751771597802085424widecontent" align="center" cellpadding="0" cellspacing="0" border="0" style="width:100%">
          <tbody><tr>
            <td bgcolor="#e0ebfa" class="m_-7751771597802085424header" style="padding:49px 0px 4px">
              
              <!-- <table align="center" border="0" cellpadding="0" cellspacing="0" style="width:100%">
                <tbody><tr>
                  <td height="70">
                    <table width="100%" border="0" cellspacing="0" cellpadding="0">
                      <tbody><tr>
                        <td align="center">
                          <a href="https://getmimo.com/" target="_blank" data-saferedirecturl="https://www.google.com/url?q=https://getmimo.com/&amp;source=gmail&amp;ust=1544385318038000&amp;usg=AFQjCNF7vZAI8EMcqexa0PP7CSGYn6n4ww">
                            <img class="m_-7751771597802085424fix CToWUd" src="https://ci3.googleusercontent.com/proxy/-4Wc1Vcjo9w5AJyh0V-YKPVf88EcLAgyvCgHUsc8o0bJojDU6aT_E8vA8ImnyJrX2SQOOQl4YbAh6_1izcXAFBIid-KOPrE3AnuBok7K8Dj0HRAD5B89=s0-d-e1-ft#https://images.mxpnl.com/1129530/1535527900835.1265385.logoimage.png" border="0" alt="" style="height:27px;width:120px">
                          </a>
                        </td>
                      </tr>
                    </tbody></table>
                  </td>
                </tr>
              </tbody></table> -->
              
            </td>
          </tr>
          <tr>
            <td bgcolor="#e0ebfa" class="m_-7751771597802085424headerimagesection" style="line-height:0px">
              
              <table align="center" border="0" cellpadding="0" cellspacing="0" style="width:100%;max-width:600px;line-height:0px">
                <tbody><tr>
                  <td>
                    <table width="100%" border="0" cellspacing="0" cellpadding="0" style="line-height:0px">
                      <tbody><tr>
                        <td align="center" style="padding:0px;vertical-align:bottom;line-height:0px">
                          <img class="m_-7751771597802085424fix CToWUd a6T" src="https://stag.uicard.io/img/mail-banner.png" border="0" alt="" style="width:100%;border-radius:10px 10px 0px 0px;height:auto" tabindex="0"><div class="a6S" dir="ltr" style="opacity: 0.01; left: 630px; top: 409px;"><div id=":12s" class="T-I J-J5-Ji aQv T-I-ax7 L3 a5q" role="button" tabindex="0" aria-label="Download attachment " data-tooltip-class="a1V" data-tooltip="Download"><div class="aSK J-J5-Ji aYr"></div></div></div>
                        </td>
                      </tr>
                    </tbody></table>
                  </td>
                </tr>
              </tbody></table>
              
            </td>
          </tr>
        </tbody></table>
        
      </td>
    </tr>
    <tr>
      <td>
        
        <table bgcolor="#ffffff" class="m_-7751771597802085424content" align="center" cellpadding="0" cellspacing="0" border="0" style="border-radius:0px 0px 10px 10px;width:100%;max-width:600px">
          <tbody><tr>
            <td class="m_-7751771597802085424innerpadding" style="padding:58px 64px 0px 62px">
              <table align="left" border="0" cellpadding="0" cellspacing="0" style="width:100%">
                <tbody><tr>
                  <td>
                    <table width="100%" border="0" cellspacing="0" cellpadding="0">
                      <tbody><tr>
                        <td class="m_-7751771597802085424h2" style="color:#3e416d;font-family:helvetica;font-size:24px;line-height:28px;font-weight:bold;padding:0 0 15px">
                          Hey there,
                        </td>
                      </tr>
                      <tr>
                        <td class="m_-7751771597802085424bodycopy" style="color:#3e416d;font-family:helvetica;font-size:18px;line-height:1.5">
                          Someone just submitted your form, here's what they had to say:
                          <br><br>
                          <!-- <hr style="margin-top:18px;margin-bottom:16px;color:#eaebf0;background-color:#eaebf0;height:1px;border:0 #eaebf0"> -->
                          
                          <!-- Real info here -->
                          <table style="font-size:16px" cellpadding="10px">
                            <tbody>` + damn + `</tbody>
                          </table>

                        </td>
                      </tr>
                        
  
                      </tr><tr>
                        <td class="m_-7751771597802085424bodycopy" style="padding-bottom:58px;color:#3e416d;font-family:helvetica;font-size:18px;line-height:1.5">
                          <br>
                          <div style="font-size:14px;text-align:center;color: #9fa0b9;">The form was submitted on ` + dateTime + ` UTC</div>
                        </td>
                      </tr>
                    </tbody></table>
                  </td>
                </tr>
              </tbody></table>
            </td>
          </tr>
				</tbody></table>
				
				<table bgcolor="#ffffff" align="center" cellpadding="0" cellspacing="0" border="0" width="100%" style="width:100%">
          <tbody><tr>
            <td class="m_-7751771597802085424footer" bgcolor="#f9fafe" style="padding:48px 30px 15px">
              <table style="width:100%" width="100%" border="0" cellspacing="0" cellpadding="0">
                <tbody><tr>
                  <td align="center" style="padding:0px 0 40px">
                    <table style="width:10%" width="10%" border="0" cellspacing="0" cellpadding="0">
                      <tbody><tr>
                        
                        <td width="40" style="padding:0 32px" align="center">
                          <a href="https://www.twitter.com/uicardio/" target="_blank">
                            <img src="https://stag.uicard.io/img/twitter.png" width="37" height="37" alt="Facebook" border="0" style="height:auto" class="CToWUd">
                          </a>
                        </td>
                        <td width="40" style="padding:0 32px" align="center">
                          <a href="https://www.instagram.com/uicardio/" target="_blank">
                            <img src="https://stag.uicard.io/img/instagram.png" width="37" height="37" alt="Instagram" border="0" style="height:auto" class="CToWUd">
                          </a>
                        </td>
                        <td width="40" style="padding:0 32px" align="center">
                          <a href="https://facebook.com/uicardio" target="_blank">
                            <img src="https://stag.uicard.io/img/facebook.png" width="37" height="37" alt="Dribble" border="0" style="height:auto" class="CToWUd">
                          </a>
                        </td>
                        
                      </tr>
                    </tbody></table>
                  </td>
                </tr>
                <tr>
                  <td align="center" style="font-family:helvetica;font-size:16px;padding:12px 0 0">
                    <div style="font-size:14px;width: 600px;margin-bottom: 30px; color: #b1b2be">
                      You are receiving this because you confirmed this email address on UICardio. If you don't remember doing that, or no longer wish to receive these emails, please remove the form on uicard.io/contact or click here to unsubscribe from this endpoint.
                    </div>
                  </td>
                </tr>
                
              </tbody></table>
            </td>
          </tr>
        </tbody></table>
        
      </td>
    </tr>
    <tr>
      <td>
        
      </td>
    </tr>
  </tbody></table>

</body>
</html>
	`
	return htmlBody
}
