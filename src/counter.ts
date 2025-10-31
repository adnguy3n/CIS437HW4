export async function updateVisitorCount(element: HTMLDivElement) {

  const cloudFunctionUrl = 'https://us-central1-cis437-hw4-476803.cloudfunctions.net/visitorCounter';

  element.innerHTML = `visitor count: ...`;

  try {
    const response = await fetch(cloudFunctionUrl, {
      method: 'GET',
    });

    if (!response.ok) {
      throw new Error(`Function call failed with status: ${response.status}`);
    }

    const data = await response.json();

    if (data && typeof data.count === 'number') {
      element.innerHTML = `visitor count: ${data.count}`;
    } else {
      throw new Error('Invalid response from function');
    }
  } catch (error) {
    console.error('Failed to fetch visitor count:', error);
    element.innerHTML = `visitor count: error`;
  }
}
