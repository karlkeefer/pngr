import {render, screen} from '@testing-library/react';
import {describe, expect, test} from 'vitest';

import SimplePage from "./SimplePage";

describe("SimplePage test", () => {
  test("Should show title", () => {
    render(<SimplePage title='Testing'><h4>Content</h4></SimplePage>);

    expect(screen.getByText(/Testing/i)).toBeDefined()
  })
})
