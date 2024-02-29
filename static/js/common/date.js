(function($) {
	$.extend({
		myTime: {
			/**
			 * <LABEL_1110>
			 * @return <int>        unix<LABEL_1511>(<LABEL_1837>)
			 */
			CurTime: function(){
				return Date.parse(new Date())/1000;
			},
			/**
			 * <LABEL_951> Unix<LABEL_1511>
			 * @param <string> 2014-01-01 20:20:20  <LABEL_1490>
			 * @return <int>        unix<LABEL_1511>(<LABEL_1837>)
			 */
			DateToUnix: function(string) {
				var f = string.split(' ', 2);
				var d = (f[0] ? f[0] : '').split('-', 3);
				var t = (f[1] ? f[1] : '').split(':', 3);
				return (new Date(
						parseInt(d[0], 10) || null,
						(parseInt(d[1], 10) || 1) - 1,
						parseInt(d[2], 10) || null,
						parseInt(t[0], 10) || null,
						parseInt(t[1], 10) || null,
						parseInt(t[2], 10) || null
					)).getTime() / 1000;
			},
			/**
			 * <LABEL_697>
			 * @param <int> unixTime    <LABEL_1491>(<LABEL_1837>)
			 * @param <bool> isFull    <LABEL_952>(Y-m-d <LABEL_1824> Y-m-d H:i:s)
			 * @param <int>  timeZone   <LABEL_1825>
			 */
			UnixToDate: function(unixTime, isFull, timeZone) {
                function add0(m){return m<10?'0'+m:m }

                if (typeof (timeZone) == 'number')
				{
					unixTime = parseInt(unixTime) + parseInt(timeZone) * 60 * 60;
				}
				var time = new Date(unixTime * 1000);
				var ymdhis = "";
				ymdhis += add0(time.getUTCFullYear()) + "-";
				ymdhis += add0((time.getUTCMonth()+1)) + "-";
				ymdhis += add0(time.getUTCDate());
				if (isFull === true)
				{
					ymdhis += " " + add0(time.getUTCHours()) + ":";
					ymdhis += add0(time.getUTCMinutes()) + ":";
					ymdhis += add0(time.getUTCSeconds());
				}
				return ymdhis;
			}
		}
	});
})(jQuery); 