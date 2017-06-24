import { MomAndIPage } from './app.po';

describe('mom-and-i App', () => {
  let page: MomAndIPage;

  beforeEach(() => {
    page = new MomAndIPage();
  });

  it('should display welcome message', () => {
    page.navigateTo();
    expect(page.getParagraphText()).toEqual('Welcome to app!!');
  });
});
