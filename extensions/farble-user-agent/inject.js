(() => {

	const USER_AGENT = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/149.0.7827.54 Safari/537.36";

	const override = (object, property, value) => {
		Object.defineProperty(object, property, {
			configurable: true,
			enumerable: true,
			get: () => value
		});
	};

	override(Navigator.prototype, "userAgent", USER_AGENT);
	override(Navigator.prototype, "appVersion", "5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/149.0.7827.54 Safari/537.36");
	override(Navigator.prototype, "platform", "Win32");
	override(Navigator.prototype, "vendor", "Google Inc.");
	override(Navigator.prototype, "productSub", "20030107");
	override(Navigator.prototype, "oscpu", undefined);

	const uaData = {
		brands: [
			{
				brand: "Google Chrome",
				version: "149"
			},
			{
				brand: "Chromium",
				version: "149"
			},
			{
				brand: "Not/A)Brand",
				version: "24"
			}
		],
		mobile: false,
		platform: "Windows",
		getHighEntropyValues: async (hints) => {

			const values = {
				architecture: "x86",
				bitness: "64",
				model: "",
				platform: "Windows",
				platformVersion: "10.0.0",
				uaFullVersion: "149.0.7827.54",
				fullVersionList: [
					{
						brand: "Google Chrome",
						version: "149.0.7827.54"
					},
					{
						brand: "Chromium",
						version: "149.0.7827.54"
					},
					{
						brand: "Not/A)Brand",
						version: "24.0.0.0"
					}
				],
				wow64: false
			};

			return Object.fromEntries(
				hints.map(hint => [hint, values[hint]])
			);

		},
		toJSON() {
			return {
				brands: this.brands,
				mobile: this.mobile,
				platform: this.platform
			};
		}
	};

	override(Navigator.prototype, "userAgentData", uaData);

})();
