function getQuoteSVG(theme = "default") {
	const URL = `https://go-stoic-quotes-production.up.railway.app/stoic-quote-svg?theme=${encodeURIComponent(
		theme
	)}`;

	// Show loading text
	$("#quote-svg").html(
		`<p class="text-gray-500 italic">Loading quote...</p>`
	);

	setTimeout(() => {
		$("#quote-svg").html(`
			<object 
				data="${URL}" 
				type="image/svg+xml" 
				width="300" 
				height="300" 
				class="mx-auto"
			></object>
		`);
	}, 200);

	// Update output preview
	$("#url").text(URL);
	$("#readme").text(`![Stoic Quotes](${URL})`);
}

$(document).ready(function () {
	getQuoteSVG();

	$("#theme").on("change", function () {
		const selectedTheme = $(this).val();
		getQuoteSVG(selectedTheme);
	});
});
