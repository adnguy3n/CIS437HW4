export async function updateVisitorCount(element: HTMLDivElement) {
  const cloudFunctionUrl = "https://us-central1-cis437-hw4-476803.cloudfunctions.net/visitorCounter";

  element.innerHTML = "visitor count: ...";

  try {
    const response = await fetch(cloudFunctionUrl);
    const data = await response.json();

    element.innerHTML = `visitor count: ${data.count}`;
  } catch {
    element.innerHTML = "visitor count: error";
  } 
}
